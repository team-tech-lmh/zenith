package main

import (
	"log"
	"os"
	"zenith/business/service"
	"zenith/zenithsdk/httpsdk"
)

func main() {
	log.SetOutput(os.Stdout)
	service.Init()
	httpsdk.StartHttpServer(":10001")
}
