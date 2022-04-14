package main

import (
	"log"
	"os"
	"zenith/business/autotask"
	"zenith/business/service"
	"zenith/zenithsdk/httpsdk"
)

func main() {
	log.SetOutput(os.Stdout)
	service.Init()
	autotask.Init()
	httpsdk.StartHttpServer(":10001")
}
