package httpsdk

import "zenith/utils"

var (
	EventKeyPrefix          = "zenithsdk-eventkey-"
	EventKeyCarPlateReceive = EventKeyPrefix + "receive-car-plate"
	EventKeyOpenBarrier     = EventKeyPrefix + "open-barrier"
	EventKeyRegisterCamera  = EventKeyPrefix + "register-camera"
)

func init() {
	utils.MessageSub(EventKeyOpenBarrier, receiveOpenBarrierEvent)
}
