package pangea

import "net/http"

type Response struct {
	OriginRes http.ResponseWriter
}

func NewResponse(res http.ResponseWriter) *Response {
	return &Response{OriginRes: res}
}
