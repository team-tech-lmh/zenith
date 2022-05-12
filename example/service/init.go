package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xizhukarsa/zenith/httpsdk"
)

func Init(e *gin.Engine) {
	registerPlateReceive()
	registerCameraFound()
	e.Any("service/barrier/open", openBarrier)
	httpsdk.SetPicSavePath("../pics/")
}
