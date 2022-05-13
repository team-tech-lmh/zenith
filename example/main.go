package main

import (
	"github.com/gin-gonic/gin"
	"github.com/team-tech-lmh/zenith/example/service"
	"github.com/team-tech-lmh/zenith/httpsdk"
)

func main() {
	e := gin.Default()
	service.Init(e)
	httpsdk.StartHttpServer(e)
	e.Run(":9091")
}
