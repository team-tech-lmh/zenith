package service

import (
	"zenith/utils"
	"zenith/zenithsdk/httpsdk"
)

func registerPlateReceive() {
	utils.MessageSub(httpsdk.EventKeyCarPlateReceive, func(msg interface{}) {
		ret := msg.(httpsdk.PlateResult)
		utils.MessagePub(httpsdk.EventKeyCarPlateReceiveCheckResult(ret.AlarmInfoPlate.IPAddr, ret.AlarmInfoPlate.Result.PlateResult.PlateID), httpsdk.PlateCheckResult{
			ShouldOpen: shouldOpenForPlate(ret.AlarmInfoPlate.Result.PlateResult.License),
		})
	})
}

func shouldOpenForPlate(licence string) bool {
	return licence == "京A20001"
}
