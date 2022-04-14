package service

import (
	"encoding/json"
	"fmt"
	"zenith/zenithsdk/httpsdk"
)

func registerPlateReceive() {
	httpsdk.SetCarPlateReceiveHandler(func(ret httpsdk.PlateResult) httpsdk.PlateCheckResult {
		return httpsdk.PlateCheckResult{
			ShouldOpen: shouldOpenForPlate(ret.AlarmInfoPlate.Result.PlateResult.License),
		}
	})
}

func registerCameraFound() {
	httpsdk.SetCameraFoundHandler(func(msg httpsdk.RegisterCameraMsg) {
		buf, _ := json.Marshal(msg)
		fmt.Printf("camera found %v\n", string(buf))
	})
}

func shouldOpenForPlate(licence string) bool {
	return licence == "京A20001"
}
