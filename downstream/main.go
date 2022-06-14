package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from downstream server for request URI: %s\n", req.RequestURI)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8081", nil)
}
