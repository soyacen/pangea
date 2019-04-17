package main

import (
	"github.com/yacen/pangea"
	"golang.org/x/exp/errors/fmt"
	"log"
	"net/http"
)

type Handler struct {
}

func (Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	res := &pangea.Response{OriginRes: response}
	req := &pangea.Request{OriginReq: request}
	//panic("implement me")
	log.Println(req.OriginReq.RequestURI)
	res.OriginRes.WriteHeader(200)
	fmt.Println(res.OriginRes.Write([]byte("hello")))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", &Handler{}))
}
