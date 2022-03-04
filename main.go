package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"zenith/zenithsdk"

	"github.com/gin-gonic/gin"
)

var (
	srv        *http.Server = nil
	openResult              = map[string]interface{}{
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
}
func baseDeferHandle(ctx *gin.Context) {
	ctx.Request.Body.Close()
	if ctx.Writer.Size() <= 0 {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	}
}

func handlePlateResult(ctx *gin.Context) {
	baseBeforeHandle(ctx)
	defer baseDeferHandle(ctx)

	buf, err := ioutil.ReadAll(ctx.Request.Body)
	if nil != err {
		log.Printf("read body failed %v\n", err)
		return
	}
	type Result struct {
		AlarmInfoPlate struct {
			Serialno string `json:"serialno"`
			IPAddr   string `json:"ipaddr"`
			Result   struct {
				PlateResult struct {
					ImageFile            string `json:"imageFile"`
					ImageFragmentFile    string `json:"imageFragmentFile"`
					ImageFileLen         int    `json:"imageFileLen"`
					ImageFragmentFileLen int    `json:"imageFragmentFileLen"`
					IsOffLine            int    `json:"isoffline"`
					License              string `json:"license"`
					PlateID              int    `json:"plateid"`
				}
			} `json:"result"`
		}
	}
	var obj Result
	if err := json.Unmarshal(buf, &obj); nil != err {
		log.Printf("parse body failed %v\n", err)
		return
	}
	if obj.AlarmInfoPlate.Result.PlateResult.License == "京AF0236" { //立即开门
		ctx.JSON(http.StatusOK, openResult)
	} else { //延迟5秒开门
		go func() {
			time.Sleep(time.Second * 5)
			openCh <- 1
		}()
	}

	go func() {
		if cli, err := zenithsdk.NewClient(obj.AlarmInfoPlate.IPAddr, 8131); nil != err {
			log.Printf("create tcp client failed %v\n", err)
		} else {
			if cmd, err := cli.ScreenShowAndSayPrice("1小时28分钟", "17元"); nil != err {
				log.Printf("show price on screen failed %v\n", err)
			} else {
				log.Printf("show price on screen result %v\n", cmd.DataString())
			}
		}
	}()
}

func handleDeviceInfo(ctx *gin.Context) {
	baseBeforeHandle(ctx)
	defer baseDeferHandle(ctx)

	// 如果有需要开门
	select {
	case <-openCh:
		ctx.JSON(http.StatusOK, openResult)
	default:
		return
	}
}

func receiveCapturedPic(ctx *gin.Context) {
	baseBeforeHandle(ctx)
	defer baseDeferHandle(ctx)

	buf, err := ioutil.ReadAll(ctx.Request.Body)
	if nil != err {
		log.Printf("read body failed %v\n", err)
		return
	}
	log.Printf("captured image content : %v\n", string(buf))
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

func startHttpServer() {
	router := gin.Default()
	router.Any("/", otherReq)
	router.Any("/devicemanagement/php/plateresult.php", handlePlateResult)
	router.Any("/devicemanagement/php/receivedeviceinfo.php", handleDeviceInfo)
	router.Any("/devicemanagement/php/receivepic.php", receiveCapturedPic)
	if err := router.Run(":10001"); nil != err {
		log.Printf("Start server failed : %s\n", err)
		panic(err)
	} else {
		log.Printf("Server start on  :10001 \n")
	}
}

func main() {
	startHttpServer()
}
