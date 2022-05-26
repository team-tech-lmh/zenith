package httpsdk

import "strings"

var (
	EventKeyPrefix = "zenithsdk-eventkey-"
)

var (
	defaultUrlConf = URLConfig{
		HeartBeat:     "/devicemanagement/php/receivedeviceinfo.php",
		ReceivePic:    "/devicemanagement/php/receivepic.php",
		ReceiveResult: "/devicemanagement/php/plateresult.php",
	}
	// 接收到车牌识别
	carPlateReceive = func(PlateResult) PlateCheckResult {
		return PlateCheckResult{
			ShouldOpen: true,
		}
	}
	// 发现摄像头设备
	cameraFound = func(RegisterCameraMsg) {
	}
	// 收到截图
	receivePic = func(CapturePic) {}
)

// url 配置
type URLConfig struct {
	HeartBeat     string
	ReceivePic    string
	ReceiveResult string
}

func SetURLPath(cnf URLConfig) {
	defaultUrlConf = cnf
}

func GetURLPath() URLConfig {
	return defaultUrlConf
}

// 开闸
type OpenBarrierPO struct {
	IPAddress string
}

func OpenBarrier(po OpenBarrierPO) {
	openBarrierAt(po.IPAddress)
}

// 图片存储位置设置
func SetPicSavePath(dirPath string) {
	picSavePath = strings.TrimRight(dirPath, "/")
}

func GetPicSavePath() string {
	return picSavePath
}

// 识别结果回调设置
func SetCarPlateReceiveHandler(f func(PlateResult) PlateCheckResult) {
	carPlateReceive = f
}

// 相机发现回调设置
func SetCameraFoundHandler(f func(RegisterCameraMsg)) {
	cameraFound = f
}

func SetTrigerPicReceiveHander(f func(CapturePic)) {
	receivePic = f
}
