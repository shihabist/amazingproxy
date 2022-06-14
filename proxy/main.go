package main

import (
	"proxy/controller"
	"proxy/repository"
	"proxy/router"
	"proxy/service"
)

var (
	fileRepo        = repository.NewFileRepository()
	proxyService    = service.NewProxyService(fileRepo)
	proxyController = controller.NewProxyController(proxyService)
	httpRouter      = router.NewStdRouter()
)

func main() {
	httpRouter.HandleReq("/", proxyController.ProxyHandler)
	httpRouter.Serve("8080")
}
