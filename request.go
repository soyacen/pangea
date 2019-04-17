package pangea

import (
	"net/http"
)

type Request struct {
	OriginReq *http.Request
	//headers http.Header
}

func NewRequest(req *http.Request) *Request {
	return &Request{OriginReq: req}
}

/*
func (req *Request) Get(name string) (string, error) {
	return req.Header(name)
}

func (req *Request) Header(name string) (string, error) {
	if name == "" {
		return "", errors.New("name argument is required to req.Header or req.Get")
	}
	return req.OriginReq.Header.Get(name), nil
}
*/
