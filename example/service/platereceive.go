package service

import (
	"encoding/json"
	"log"

	"github.com/team-tech-lmh/zenith/httpsdk"
)

func registerPlateReceive() {
	httpsdk.SetCarPlateReceiveHandler(func(ret httpsdk.PlateResult) httpsdk.PlateCheckResult {
		buf, _ := json.Marshal(ret)
		log.Printf("plate found %v\n", string(buf))
		return httpsdk.PlateCheckResult{
			ShouldOpen:   shouldOpenForPlate(ret.AlarmInfoPlate.Result.PlateResult.License),
			VoiceContent: "月租车",
		}
	})
}

func registerCameraFound() {
	httpsdk.SetCameraFoundHandler(func(msg httpsdk.RegisterCameraMsg) {
		buf, _ := json.Marshal(msg)
		log.Printf("camera found %v\n", string(buf))
	})
}

func shouldOpenForPlate(licence string) bool {
	return licence == "京AF0236"
}
