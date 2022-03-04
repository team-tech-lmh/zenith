package httpsdk

import (
	"fmt"
	"zenith/utils"
)

var (
	EventKeyPrefix                     = "zenithsdk-eventkey-"
	EventKeyCarPlateReceive            = EventKeyPrefix + "receive-car-plate"
	EventKeyCarPlateReceiveCheckResult = func(ipAddr string, plated int) string {
		return EventKeyPrefix + "receive-car-plate-check-result-" + ipAddr + fmt.Sprintf("-%v-", plated)
	}
	EventKeyOpenBarrier    = EventKeyPrefix + "open-barrier"
	EventKeyRegisterCamera = EventKeyPrefix + "register-camera"
)

func init() {
	utils.MessageSub(EventKeyOpenBarrier, receiveOpenBarrierEvent)
}
