package service

import (
	"encoding/json"

	"github.com/team-tech-lmh/zenith/httpsdk"
	"github.com/team-tech-lmh/zenith/utils"
)

func registerPlateReceive() {
	httpsdk.SetCarPlateReceiveHandler(func(ret httpsdk.PlateResult) httpsdk.PlateCheckResult {
		buf, _ := json.Marshal(ret)
		utils.DefaultSwitchLogger.Printf("plate found %v\n", string(buf))
		return httpsdk.PlateCheckResult{
			ShouldOpen:   shouldOpenForPlate(ret.AlarmInfoPlate.Result.PlateResult.License),
			VoiceContent: "月租车",
		}
	})
}

func registerCameraFound() {
	httpsdk.SetCameraFoundHandler(func(msg httpsdk.RegisterCameraMsg) {
		buf, _ := json.Marshal(msg)
		utils.DefaultSwitchLogger.Printf("camera found %v\n", string(buf))
	})
}

func shouldOpenForPlate(licence string) bool {
	return licence == "京AF0236"
}
