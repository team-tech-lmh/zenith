package httpsdk

type RegisterCameraMsg struct {
	IPAddr string
}

func registerCamera(ipAddr string) {
	cameraFound(RegisterCameraMsg{
		IPAddr: ipAddr,
	})
}
