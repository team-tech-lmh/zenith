package main

import "zenith/zenithsdk/httpsdk"

func main() {
	httpsdk.StartHttpServer(":10001")
}
