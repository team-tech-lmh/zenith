package main

import (
	"github.com/xizhukarsa/zenith/example/service"
	"github.com/xizhukarsa/zenith/httpsdk"
)

func main() {
	service.Init()
	httpsdk.StartHttpServer(":10001")
}
