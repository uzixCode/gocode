package models

type BaseRequest struct {
	Path   string      `Json:"path"`
	Method string      `Json:"method"`
	Data   interface{} `Json:"data"`
}
