package websocket

import (
	"EX_binancequant/mylog"
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
	// 启动一个gouroutine发送心跳
	go func() {
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
	}()

	// 这是一个同步处理模型（只是一个例子），如果希望并行处理可以每个请求一个gorutine，注意控制并发goroutine的数量!!!
	for {
		if wsConn.isClosed {
			break
		}

		msg, err := wsConn.wsRead()
		if err != nil {
			mylog.DataLogger.Error().Msgf("[Websocket] read message fail err: %v", err)
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
		jsonResponse := new(JsonResponse)
		jsonRequest.Method = j.Get("method").MustString()
		jsonRequest.Symbol = j.Get("symbol").MustString()
		jsonRequest.Interval = j.Get("interval").MustString()
		jsonRequest.Levels = j.Get("levels").MustString()
		jsonRequest.ListenKey = j.Get("listenKey").MustString()
		jsonRequest.ID = j.Get("id").MustInt64()

		jsonResponse.Result = ""
		jsonResponse.ID = jsonRequest.ID
		response, err := json.Marshal(jsonResponse)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[Websocket] json Marshal fail err: %v", err)
			wsConn.wsClose()
			break
		}

		switch jsonRequest.Method {
		case "normal":
			// 推送BasicMessageData
			go PushAllMarkPrice(wsConn)
			go PushKline(wsConn, jsonRequest.Symbol, jsonRequest.Interval)
			go PushAggTrade(wsConn, jsonRequest.Symbol)
			go PushAllMarketsStat(wsConn)
			go PushDepth(wsConn, jsonRequest.Symbol)
			go PushUserData(wsConn, jsonRequest.ListenKey)
		case "kline":
			// 推送KlineMessageData
			go PushKlineInterval(wsConn, jsonRequest.Symbol, jsonRequest.Interval)
		case "depth":
			// 推送KlineMessageData
			go PushDepthLevels(wsConn, jsonRequest.Symbol, jsonRequest.Levels)
		default:
			mylog.DataLogger.Error().Msgf("[Websocket] read message param err")
			jsonResponse.Result = "request message param err"
		}

		err = wsConn.wsWrite(msg.messageType, response)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[Websocket] write message fail err: %v", err)
			wsConn.wsClose()
			break
		}
	}
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
