package service

import "zenith/zenithsdk/httpsdk"

func Init() {
	registerPlateReceive()
	registerCameraFound()
	httpsdk.SetPicSavePath("/Users/karsa/Downloads")
	httpsdk.RegisterAPI("service/barrier/open", openBarrier)
}
