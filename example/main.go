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
	res := pangea.NewResponse(response)
	req := pangea.NewRequest(request)
	log.Println(req.OriginReq.RequestURI)
	res.OriginRes.WriteHeader(200)
	fmt.Println(res.OriginRes.Write([]byte("hello")))
}

func main() {

	app := pangea.Application{}

	server := http.Server{
		Addr:    ":8080",
		Handler: &Handler{},
	}
	log.Fatal(server.ListenAndServe())
}
