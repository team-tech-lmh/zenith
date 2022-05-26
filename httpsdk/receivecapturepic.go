package httpsdk

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/team-tech-lmh/zenith/utils"
)

type CapturePic struct {
	TriggerImage struct {
		ImageFile          string `json:"imageFile"`
		ImageFileBase64Len int    `json:"imageFileBase64Len"`
		ImageFileLen       int    `json:"imageFileLen"`
	}
}

func (ret CapturePic) SavePic() (map[PicType]string, error) {
	filePath, err := saveCatpurePicBase64Content(PicTypeTrigger, ret.TriggerImage.ImageFile)
	if nil != err {
		return nil, err
	}
	return map[PicType]string{
		PicTypeTrigger: filePath,
	}, nil
}

func receiveCapturedPic(ctx *gin.Context) {
	baseBeforeHandle(ctx)
	defer baseDeferHandle(ctx)

	buf, err := ioutil.ReadAll(ctx.Request.Body)
	if nil != err {
		utils.DefaultSwitchLogger.Printf("read body failed %v\n", err)
		return
	}
	utils.DefaultSwitchLogger.Printf("receiveCapturedPic -- %v\n", string(buf))

	var ret CapturePic
	if err := json.Unmarshal(buf, &ret); nil != err {
		utils.DefaultSwitchLogger.Printf("receiveCapturedPic parse body failed %v\n", err)
		return
	}
	receivePic(ret)
}
