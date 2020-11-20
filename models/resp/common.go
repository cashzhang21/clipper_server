package resp

// Response
type Response struct {
	Code    int         `json:"errorno"`
	Message string      `json:"errormsg"`
	Data    interface{} `json:"data"`
}
