package httpsdk

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	gRouter    *gin.Engine = nil
	openResult             = map[string]interface{}{
		"Response_AlarmInfoPlate": map[string]interface{}{
			"info": "ok",
			"TriggerImage": map[string]interface{}{
				"port":                 10001,
				"snapImageRelativeUrl": defaultUrlConf.ReceivePic,
			},
		},
	}
)

func baseBeforeHandle(ctx *gin.Context) {
	log.Printf("url %v\n", ctx.Request.URL.Path)
	ip, has := ctx.RemoteIP()
	if has {
		remoteAddrFind(ip.String())
	}
}

func baseDeferHandle(ctx *gin.Context) {
	ctx.Request.Body.Close()
	if ctx.Writer.Size() <= 0 {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	}
}

func remoteAddrFind(ipAddr string) {
	log.Printf("ipaddr found %v\n ", ipAddr)
	registerCamera(ipAddr)
}

func StartHttpServer(router *gin.Engine) {
	gRouter = router
	router.Any(defaultUrlConf.ReceiveResult, handlePlateResult)
	router.Any(defaultUrlConf.HeartBeat, handleHeartBeat)
	router.Any(defaultUrlConf.ReceivePic, receiveCapturedPic)
}
