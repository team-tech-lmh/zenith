package httpsdk

var (
	EventKeyPrefix = "zenithsdk-eventkey-"
)
var (
	// 接收到车牌识别
	carPlateReceive = func(PlateResult) PlateCheckResult {
		return PlateCheckResult{
			ShouldOpen: true,
		}
	}
	// 发现摄像头设备
	cameraFound = func(RegisterCameraMsg) {

	}
)

func SetCarPlateReceiveHandler(f func(PlateResult) PlateCheckResult) {
	carPlateReceive = f
}

func SetCameraFoundHandler(f func(RegisterCameraMsg)) {
	cameraFound = f
}
