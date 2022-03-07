package httpsdk

type OpenBarrierPO struct {
	IPAddress string
}

func OpenBarrier(po OpenBarrierPO) {
	openBarrierAt(po.IPAddress)
}
