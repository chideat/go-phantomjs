package phantomjs

type Response struct {
	Status    int    `json:"status"`
	SessionId string `json:"sessionId"`
}
