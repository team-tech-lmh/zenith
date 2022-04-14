package service

import (
	"zenith/zenithsdk/httpsdk"
)

func registerPlateReceive() {
	httpsdk.SetCarPlateReceiveHandler(func(ret httpsdk.PlateResult) httpsdk.PlateCheckResult {
		return httpsdk.PlateCheckResult{
			ShouldOpen: shouldOpenForPlate(ret.AlarmInfoPlate.Result.PlateResult.License),
		}
	})
}

func shouldOpenForPlate(licence string) bool {
	return licence == "京A20001"
}
