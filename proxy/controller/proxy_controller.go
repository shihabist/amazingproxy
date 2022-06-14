package controller

import (
	// "crypto/tls"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"net/url"

	"proxy/errors"
	"proxy/model"
	"proxy/service"
)

type ProxyController interface {
	GetProxy(writer http.ResponseWriter, request *http.Request)
}
type controller struct{}

var (
	proxyService service.ProxyService
)

func NewProxyController(service service.ProxyService) ProxyController {
	proxyService = service
	return &controller{}
}
func (*controller) GetProxy(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")
	incomingUri := model.AllowedUri{
		Uri:    request.RequestURI,
		Method: request.Method,
	}
	isAllowed := proxyService.Validate(&incomingUri)
	if !isAllowed {
		writer.WriteHeader(http.StatusUnauthorized)
		if err := json.NewEncoder(writer).Encode(errors.ServiceError{Message: "failed to get employees"}); err != nil {
			return
		}
		return
	}
	targetUrl, err := url.Parse("http://172.21.0.6:8081")
	if err != nil {
		panic(err)
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	httputil.NewSingleHostReverseProxy(targetUrl).ServeHTTP(writer, request)
}
