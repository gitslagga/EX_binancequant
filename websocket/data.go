package websocket

type JsonRequest struct {
	Method    string `json:"method"`
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	Levels    string `json:"levels"`
	ListenKey string `json:"listenKey"`
	ID        int64  `json:"id"`
}

type JsonResponse struct {
	Result string `json:"result"`
	ID     int64  `json:"id"`
}
