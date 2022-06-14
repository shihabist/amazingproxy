package router

import (
	"fmt"
	"log"
	"net/http"
	"time"

	logger "proxy/log"
)

type stdRouter struct{}

func NewStdRouter() Router {
	return &stdRouter{}
}

var (
	stdRouterDispatch = http.NewServeMux()
)

func (*stdRouter) HandleReq(uri string, proxyHandler func(writer http.ResponseWriter, request *http.Request)) {
	stdRouterDispatch.HandleFunc(uri, proxyHandler)
}

func (*stdRouter) Serve(port string) {
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      stdRouterDispatch,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	server.SetKeepAlivesEnabled(true)
	logger.Print(fmt.Sprintf("Listening on %s\n", server.Addr), "info")
	log.Fatal(server.ListenAndServe())
}
