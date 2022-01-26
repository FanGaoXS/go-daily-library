package response

import "time"

type Response struct {
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
	Code      int64       `json:"code"`
	Message   string      `json:"message"`
}

func New() *Response {
	return &Response{
		Data:      nil,
		Timestamp: time.Now().Unix(),
		Code:      2000,
	}
}

func (this *Response) SetMessage(message string) *Response {
	this.Message = message
	return this
}

func (this *Response) SetData(data interface{}) *Response {
	this.Data = data
	return this
}

func (this *Response) SetCode(code int64) *Response {
	this.Code = code
	return this
}
