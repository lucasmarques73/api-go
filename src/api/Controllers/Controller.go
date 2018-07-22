package Controllers

// Response - Controller Response Struct
type Response struct {
	Errors  bool        `json:"errors"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
