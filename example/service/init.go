package service

import "zenith/zenithsdk/httpsdk"

func Init() {
	registerPlateReceive()
	registerCameraFound()
	httpsdk.SetPicSavePath("../pics/")
	httpsdk.RegisterAPI("service/barrier/open", openBarrier)
}
