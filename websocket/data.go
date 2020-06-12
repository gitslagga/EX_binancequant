package websocket

type JsonRequest struct {
	Method string `json:"method"`
	Params string `json:"params"`
	ID     int64  `json:"id"`
}
