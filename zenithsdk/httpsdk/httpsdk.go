package httpsdk

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	router     = gin.Default()
	openResult = map[string]interface{}{
		"Response_AlarmInfoPlate": map[string]interface{}{
			"info": "ok",
			"TriggerImage": map[string]interface{}{
				"port":                 10001,
				"snapImageRelativeUrl": "/devicemanagement/php/receivepic.php",
			},
		},
	}
	openCh = make(chan int, 1024)
)

func baseBeforeHandle(ctx *gin.Context) {
	log.Printf("url %v\n", ctx.Request.URL.Path)
	if ctx.Request.URL.Path == "/devicemanagement/php/receivedeviceinfo1.php" {
		log.Printf("url %v\n", ctx.Request.URL.Path)
	}
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

func otherReq(ctx *gin.Context) {
	baseBeforeHandle(ctx)
	defer baseDeferHandle(ctx)

	buf, err := ioutil.ReadAll(ctx.Request.Body)
	if nil != err {
		log.Printf("read body failed %v\n", err)
		return
	}
	log.Printf("other request %v\n", string(buf))
}

func remoteAddrFind(ipAddr string) {
	log.Printf("ipaddr found %v\n ", ipAddr)
	registerCamera(ipAddr)
}

func RegisterAPI(path string, f gin.HandlerFunc) {
	router.Any(path, f)
}

func StartHttpServer(addr string) {
	router.Any("/", otherReq)
	router.Any("/websocket", webSocket)
	router.Any("/devicemanagement/php/plateresult.php", handlePlateResult)
	router.Any("/devicemanagement/php/receivedeviceinfo.php", handleHeartBeat)
	router.Any("/devicemanagement/php/receivepic.php", receiveCapturedPic)
	if err := router.Run(addr); nil != err {
		log.Printf("Start server failed : %s\n", err)
		panic(err)
	} else {
		log.Printf("Server start on  :10001 \n")
	}
}
