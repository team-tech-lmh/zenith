package service

import "zenith/httpsdk"

func Init() {
	registerPlateReceive()
	registerCameraFound()
	httpsdk.SetPicSavePath("../pics/")
	httpsdk.RegisterAPI("service/barrier/open", openBarrier)
}
