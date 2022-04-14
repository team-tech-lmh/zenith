package httpsdk

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"zenith/zenithsdk/tcpsdk"

	"github.com/gin-gonic/gin"
)

type PlateCheckResult struct {
	ShouldOpen bool
}

type PlateResult struct {
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

func handlePlateResult(ctx *gin.Context) {
	baseBeforeHandle(ctx)
	defer baseDeferHandle(ctx)

	buf, err := ioutil.ReadAll(ctx.Request.Body)
	if nil != err {
		log.Printf("read body failed %v\n", err)
		return
	}

	var obj PlateResult
	if err := json.Unmarshal(buf, &obj); nil != err {
		log.Printf("parse body failed %v\n", err)
		return
	}

	go saveCatpurePicBase64Content(PicTypeRecognizeFile, obj.AlarmInfoPlate.Result.PlateResult.ImageFile)
	go saveCatpurePicBase64Content(PicTypeRecognizeFragmentFile, obj.AlarmInfoPlate.Result.PlateResult.ImageFragmentFile)

	ret := carPlateReceive(obj)
	if !ret.ShouldOpen {
		return
	}

	ctx.JSON(http.StatusOK, openResult)
	go func() {
		if cli, err := tcpsdk.NewClient(obj.AlarmInfoPlate.IPAddr, 8131); nil != err {
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
