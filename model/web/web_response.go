package web

type WebResponse struct {
	Code   uint16      `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
