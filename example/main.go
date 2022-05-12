package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xizhukarsa/zenith/example/service"
	"github.com/xizhukarsa/zenith/httpsdk"
)

func main() {
	e := gin.Default()		
	service.Init(e)
	httpsdk.StartHttpServer(e)
	e.Run(":9091")
}
