package zenithsdk

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var (
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
	remoteAddrFind(ctx.Request.RemoteAddr)
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
	log.Printf("receive reslt %v\n", string(buf))

	if obj.AlarmInfoPlate.Result.PlateResult.License == "京AF0236" { //立即开门
		ctx.JSON(http.StatusOK, openResult)
	} else { //延迟5秒开门
		go func() {
			time.Sleep(time.Second * 5)
			openCh <- 1
		}()
	}

	go func() {
		if cli, err := NewClient(obj.AlarmInfoPlate.IPAddr, 8131); nil != err {
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
	type CapturePic struct {
		ImageFile          string `json:"imageFile"`
		ImageFileBase64Len int    `json:"imageFileBase64Len"`
		ImageFileLen       int    `json:"imageFileLen"`
	}
	var ret CapturePic
	if err := json.Unmarshal(buf, &ret); nil != err {
		log.Printf("receiveCapturedPic parse body failed %v\n", err)
		return
	}
	log.Printf("save file with len %v base64 len %v\n", ret.ImageFileLen, ret.ImageFileBase64Len)
	go catpurePic(ret.ImageFile)

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
}

func catpurePic(picContent string) {
	t := time.Now().Unix()
	md5Byte := md5.Sum([]byte(picContent))
	md5Str := base64.StdEncoding.EncodeToString(md5Byte[:])
	fName := fmt.Sprintf("%v-%v.png", md5Str, t)
	f, err := os.Create(fmt.Sprintf("/Users/karsa/Desktop/%v", fName))
	if nil != err {
		log.Printf("save pic failed when create file %v\n", err)
		return
	}
	defer f.Close()

	buf, err := base64.StdEncoding.DecodeString(picContent)
	if nil != err {
		log.Printf("save pic failed when decode file content %v\n", err)
		return
	}
	if l, err := f.Write(buf); nil != err {
		log.Printf("save pic failed when write file content %v\n", err)
		return
	} else {
		log.Printf("save pic to %v (len:%v) \n", fName, l)
	}
}

func StartHttpServer(addr string) {
	router := gin.Default()
	router.Any("/", otherReq)
	router.Any("/devicemanagement/php/plateresult.php", handlePlateResult)
	router.Any("/devicemanagement/php/receivedeviceinfo.php", handleDeviceInfo)
	router.Any("/devicemanagement/php/receivepic.php", receiveCapturedPic)
	if err := router.Run(addr); nil != err {
		log.Printf("Start server failed : %s\n", err)
		panic(err)
	} else {
		log.Printf("Server start on  :10001 \n")
	}
}
