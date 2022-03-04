package service

import "zenith/zenithsdk/httpsdk"

func Init() {
	registerPlateReceive()
	httpsdk.RegisterAPI("service/barrier/open", openBarrier)
}
