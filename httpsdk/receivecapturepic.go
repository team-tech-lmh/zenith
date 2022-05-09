package httpsdk

import (
	"encoding/json"
	"io/ioutil"
	"log"

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
	go saveCatpurePicBase64Content(PicTypeTrigger, ret.TriggerImage.ImageFile)
}
