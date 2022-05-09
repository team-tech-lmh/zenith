package main

import (
	"zenith/example/service"
	"zenith/httpsdk"
)

func main() {
	service.Init()
	httpsdk.StartHttpServer(":10001")
}
