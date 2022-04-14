package service

import "zenith/zenithsdk/httpsdk"

func Init() {
	registerPlateReceive()
	registerCameraFound()
	httpsdk.RegisterAPI("service/barrier/open", openBarrier)
}
