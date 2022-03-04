package httpsdk

type OpenBarrierEvent struct {
	IPAddress string
}

func receiveOpenBarrierEvent(msg interface{}) {
	openBarrierAt(msg.(OpenBarrierEvent).IPAddress)
}
