package main

import (
	"zenith/example/service"
	"zenith/zenithsdk/httpsdk"
)

func main() {
	service.Init()
	httpsdk.StartHttpServer(":10001")
}