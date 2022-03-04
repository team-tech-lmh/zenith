package httpsdk

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"zenith/zenithsdk/tcpsdk"

	"github.com/gin-gonic/gin"
)

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
