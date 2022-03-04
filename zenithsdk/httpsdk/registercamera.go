package httpsdk

import "zenith/utils"

type RegisterCameraMsg struct {
	IPAddr string
}

func registerCamera(ipAddr string) {
	utils.MessagePub(EventKeyRegisterCamera, RegisterCameraMsg{
		IPAddr: ipAddr,
	})
}
