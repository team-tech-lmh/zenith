package httpsdk

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func receiveCapturedPic(ctx *gin.Context) {
	baseBeforeHandle(ctx)
	defer baseDeferHandle(ctx)

	buf, err := ioutil.ReadAll(ctx.Request.Body)
	if nil != err {
		log.Printf("read body failed %v\n", err)
		return
	}
	log.Printf("receiveCapturedPic -- %v\n", string(buf))
	type CapturePic struct {
		TriggerImage struct {
			ImageFile          string `json:"imageFile"`
			ImageFileBase64Len int    `json:"imageFileBase64Len"`
			ImageFileLen       int    `json:"imageFileLen"`
		}
	}
	var ret CapturePic
	if err := json.Unmarshal(buf, &ret); nil != err {
		log.Printf("receiveCapturedPic parse body failed %v\n", err)
		return
	}
	log.Printf("save file with len %v base64 len %v\n", ret.TriggerImage.ImageFileLen, ret.TriggerImage.ImageFileBase64Len)
	go saveCatpurePic(ret.TriggerImage.ImageFile)
}

func saveCatpurePic(picContent string) {
	if len(picContent) <= 0 {
		return
	}
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
