package websocket

import (
	"EX_binancequant/mylog"
	"EX_binancequant/trade/futures"
	"encoding/json"
	"errors"
	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

// http升级websocket协议的配置
var wsUpgrader = websocket.Upgrader{
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 客户端读写消息
type wsMessage struct {
	messageType int
	data        []byte
}

// 客户端连接
type wsConnection struct {
	wsSocket *websocket.Conn // 底层websocket
	inChan   chan *wsMessage // 读队列
	outChan  chan *wsMessage // 写队列

	mutex     sync.Mutex // 避免重复关闭管道
	isClosed  bool
	closeChan chan byte // 关闭通知
}

//数据发送类型
var MessageType = websocket.TextMessage

func (wsConn *wsConnection) wsReadLoop() {
	for {
		// 读一个message
		msgType, data, err := wsConn.wsSocket.ReadMessage()
		if err != nil {
			goto error
		}
		req := &wsMessage{
			msgType,
			data,
		}
		// 放入请求队列
		select {
		case wsConn.inChan <- req:
		case <-wsConn.closeChan:
			goto closed
		}
	}
error:
	wsConn.wsClose()
closed:
}

func (wsConn *wsConnection) wsWriteLoop() {
	for {
		select {
		// 取一个应答
		case msg := <-wsConn.outChan:
			// 写给websocket
			if err := wsConn.wsSocket.WriteMessage(msg.messageType, msg.data); err != nil {
				goto error
			}
		case <-wsConn.closeChan:
			goto closed
		}
	}
error:
	wsConn.wsClose()
closed:
}

func (wsConn *wsConnection) procLoop() {
	go dataHandler(wsConn)
}

func WSHandler(resp http.ResponseWriter, req *http.Request) {
	// 应答客户端告知升级连接为websocket
	wsSocket, err := wsUpgrader.Upgrade(resp, req, nil)
	if err != nil {
		return
	}
	wsConn := &wsConnection{
		wsSocket:  wsSocket,
		inChan:    make(chan *wsMessage, 1000),
		outChan:   make(chan *wsMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
	}

	// 处理器
	go wsConn.procLoop()
	// 读协程
	go wsConn.wsReadLoop()
	// 写协程
	go wsConn.wsWriteLoop()
}

func (wsConn *wsConnection) wsWrite(messageType int, data []byte) error {
	select {
	case wsConn.outChan <- &wsMessage{messageType, data}:
	case <-wsConn.closeChan:
		return errors.New("websocket closed")
	}
	return nil
}

func (wsConn *wsConnection) wsRead() (*wsMessage, error) {
	select {
	case msg := <-wsConn.inChan:
		return msg, nil
	case <-wsConn.closeChan:
	}
	return nil, errors.New("websocket closed")
}

func (wsConn *wsConnection) wsClose() {
	wsConn.wsSocket.Close()

	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if !wsConn.isClosed {
		wsConn.isClosed = true
		close(wsConn.closeChan)
	}
}

/**
启动一个gouroutine发送心跳
*/
func heartBeat(wsConn *wsConnection) {
	for {
		if wsConn.isClosed {
			break
		}

		time.Sleep(60 * time.Second)

		jsonResponse := new(JsonResponse)
		jsonResponse.Result = "pong"
		jsonResponse.ID = 0
		response, err := json.Marshal(jsonResponse)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[Websocket] json Marshal fail err: %v", err)
			wsConn.wsClose()
			break
		}

		if err := wsConn.wsWrite(MessageType, response); err != nil {
			mylog.DataLogger.Error().Msgf("[Websocket] pong write fail err: %v", err)
			wsConn.wsClose()
			break
		}
	}
}

/**
注意控制并发goroutine的数量!!!
*/
func dataHandler(wsConn *wsConnection) {
	for {
		if wsConn.isClosed {
			break
		}

		msg, err := wsConn.wsRead()
		if err != nil {
			//mylog.DataLogger.Error().Msgf("[Websocket] read message fail err: %v", err)
			wsConn.wsClose()
			break
		}

		mylog.DataLogger.Info().Msgf("[Websocket] read message data: %v", string(msg.data))

		j, err := simplejson.NewJson(msg.data)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[Websocket] read message fail err: %v", err)
			wsConn.wsClose()
			break
		}

		jsonRequest := new(JsonRequest)
		jsonRequest.ID = j.Get("id").MustInt64()
		jsonRequest.Method = j.Get("method").MustString()
		jsonRequest.Params = j.Get("params").MustStringArray()

		// 推送币安交易数据
		go wsConn.PushTradeData(jsonRequest)
	}
}

func (wsConn *wsConnection) PushTradeData(jsonRequest *JsonRequest) {
	wsDepthHandler := func(event []byte) {
		if !wsConn.isClosed {
			err := wsConn.wsWrite(MessageType, event)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushTradeData] write message fail err: %v", err)
			}
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[PushTradeData] WsCombinedTradeDataServe handler fail err: %v", err)
	}

	_, stopC, err := futures.WsCombinedTradeDataServe(jsonRequest.Params, wsDepthHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[PushTradeData] WsCombinedTradeDataServe dial fail err: %v", err)
		return
	}

	for {
		time.Sleep(60 * time.Second)
		if wsConn.isClosed {
			stopC <- struct{}{}
			break
		}
	}
}
