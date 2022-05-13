package service

import (
	"encoding/json"
	"fmt"

	"github.com/team-tech-lmh/zenith/httpsdk"
)

func registerPlateReceive() {
	httpsdk.SetCarPlateReceiveHandler(func(ret httpsdk.PlateResult) httpsdk.PlateCheckResult {
		buf, _ := json.Marshal(ret)
		fmt.Printf("plate found %v\n", string(buf))
		return httpsdk.PlateCheckResult{
			ShouldOpen:   shouldOpenForPlate(ret.AlarmInfoPlate.Result.PlateResult.License),
			VoiceContent: "月租车",
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
	return licence == "京AF0236"
}
