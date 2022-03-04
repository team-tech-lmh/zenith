package main

import (
	"zenith/business/service"
	"zenith/zenithsdk/httpsdk"
)

func main() {
	service.Init()
	httpsdk.StartHttpServer(":10001")
}
