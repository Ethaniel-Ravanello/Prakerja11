package base

type BaseResponse struct {
	Error   bool
	Code    int
	Message string
	Data    interface{}
}
