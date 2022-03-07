package httpsdk

var (
	EventKeyPrefix = "zenithsdk-eventkey-"

	EventKeyRegisterCamera = EventKeyPrefix + "register-camera"
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
