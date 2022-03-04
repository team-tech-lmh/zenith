package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"zenith/zenithsdk"

	"github.com/gin-gonic/gin"
)

var (
	srv        *http.Server = nil
	openResult              = map[string]interface{}{
		"Response_AlarmInfoPlate": map[string]interface{}{
			"info":         "ok",
			"is_pay":       "true",
			"content":      "怎么说",
			"TriggerImage": map[string]interface{}{},
		},
	}
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

	go func() {
		if cli, err := zenithsdk.NewClient("192.168.2.233", 8131); nil != err {
			log.Printf("create tcp client failed %v\n", err)
		} else {
			if cmd, err := cli.ScreenShowAndSayPrice("1小时28分钟", "17元"); nil != err {
				log.Printf("show price on screen failed %v\n", err)
			} else {
				log.Printf("show price on screen result %v\n", cmd.DataString())
			}
		}
	}()

	// ctx.JSON(http.StatusOK, openResult)
}
func handleDeviceInfo(ctx *gin.Context) {
	baseBeforeHandle(ctx)
	defer baseDeferHandle(ctx)
	buf, _ := ioutil.ReadAll(ctx.Request.Body)
	fmt.Printf("body %v\n", string(buf))

	// form, err := ctx.MultipartForm()
	// if nil != err {
	// 	log.Printf("read form failed %v\n", err)
	// 	return
	// }
	// if len(form.Value) > 0 {
	// 	buf, err := json.Marshal(form.Value)
	// 	if nil != err {
	// 		log.Printf("read form failed %v\n", err)
	// 		return
	// 	}
	// 	log.Println(string(buf))
	// } else if len(form.File) > 0 {
	// 	for k, fList := range form.File {
	// 		for _, f := range fList {
	// 			log.Printf("%v - %v\n", k, f.Filename)
	// 		}
	// 	}
	// } else {
	// 	buf, _ := ioutil.ReadAll(ctx.Request.Body)
	// 	fmt.Printf("body %v\n", string(buf))
	// }

	// ret := map[string]interface{}{
	// 	"Response_AlarmInfoPlate": map[string]interface{}{
	// 		"info":   "ok",
	// 		"is_pay": "true",
	// 		"TriggerImage": map[string]interface{}{
	// 			"port":                 10001,
	// 			"snapImageRelativeUrl": "/devicemanagement/php/receivedeviceinfo1.php",
	// 		},
	// 	},
	// }
	// ctx.JSON(http.StatusOK, ret)
}
func otherReq(ctx *gin.Context) {
	baseBeforeHandle(ctx)
	defer baseDeferHandle(ctx)

	buf, err := ioutil.ReadAll(ctx.Request.Body)
	if nil != err {
		log.Printf("read body failed %v\n", err)
		return
	}
	log.Println(string(buf))
}

func startHttpServer() {
	router := gin.Default()
	router.Any("/", otherReq)
	router.Any("/devicemanagement/php/plateresult.php", handlePlateResult)
	router.Any("/devicemanagement/php/receivedeviceinfo.php", handleDeviceInfo)
	router.Any("/devicemanagement/php/receivedeviceinfo1.php", otherReq)
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
