package httpsdk

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/team-tech-lmh/zenith/tcpsdk"
	"github.com/team-tech-lmh/zenith/utils"

	"github.com/gin-gonic/gin"
)

type PlateCheckResult struct {
	ShouldOpen   bool
	VoiceContent string
}

type PlateType int

const (
	PlateTypeUnknow                             PlateType = iota
	PlateTypeBlue                                         //蓝牌
	PlateTypeBlack                                        //黑牌
	PlateTypeSingleLineYellow                             //单排黄牌
	PlateTypeDoubleLineYellow                             //双排黄牌
	PlateTypePolice                                       //警车车牌
	PlateTypeArmedPolice                                  //武警车牌
	PlateTypeCustom                                       //个性化车牌PlateTypeArmedPolicePlate
	PlateTypeSingleLineArmy                               //单排军车车牌
	PlateTypeDoubleLineArmy                               //双排军车车牌
	PlateTypeConsulate                                    //领事馆车牌
	PlateTypeHongKongToMainLand                           //香港进出中国大陆车牌
	PlateTypeAgriculturalVehicle                          //农用车车牌
	PlateTypeCoachCar                                     //教练车牌
	PlateTypeMacaoToMainLand                              //澳门进出中国大陆车牌
	PlateTypeDoubleLineArmedPolice                        //双层武警车牌
	PlateTypeHeadquarterOfArmedPolice                     //武警总队车牌
	PlateTypeDoubleLineHeadquarterOfArmedPolice           //双层武警总队车牌
	PlateTypeCivilAviation                                //民航车牌
	PlateTypeNewEnergy                                    //新能源车牌

)

type PlateResult struct {
	AlarmInfoPlate struct {
		Serialno string `json:"serialno"`
		IPAddr   string `json:"ipaddr"`
		Result   struct {
			PlateResult struct {
				ImageFile            string    `json:"imageFile"`
				ImageFragmentFile    string    `json:"imageFragmentFile"`
				ImageFileLen         int       `json:"imageFileLen"`
				ImageFragmentFileLen int       `json:"imageFragmentFileLen"`
				IsOffLine            int       `json:"isoffline"`
				License              string    `json:"license"`
				PlateID              int       `json:"plateid"`
				Type                 PlateType `json:"type"`
			}
		} `json:"result"`
	}
}

func (ret PlateResult) SavePic() (map[PicType]string, error) {
	filePath1, err := saveCatpurePicBase64Content(PicTypeRecognizeFile, ret.AlarmInfoPlate.Result.PlateResult.ImageFile)
	if nil != err {
		return nil, err
	}
	filePath2, err := saveCatpurePicBase64Content(PicTypeRecognizeFragmentFile, ret.AlarmInfoPlate.Result.PlateResult.ImageFragmentFile)
	if nil != err {
		return nil, err
	}
	return map[PicType]string{
		PicTypeRecognizeFile:         filePath1,
		PicTypeRecognizeFragmentFile: filePath2,
	}, nil
}

func handlePlateResult(ctx *gin.Context) {
	baseBeforeHandle(ctx)
	defer baseDeferHandle(ctx)

	buf, err := ioutil.ReadAll(ctx.Request.Body)
	if nil != err {
		utils.DefaultSwitchLogger.Printf("read body failed %v\n", err)
		return
	}

	utils.DefaultSwitchLogger.Println("plate result ---------- " + string(buf))

	var obj PlateResult
	if err := json.Unmarshal(buf, &obj); nil != err {
		utils.DefaultSwitchLogger.Printf("parse body failed %v\n", err)
		return
	}

	ret := carPlateReceive(obj)
	if !ret.ShouldOpen {
		return
	}

	ctx.JSON(http.StatusOK, openResult)

	go func() {
		if cli, err := tcpsdk.NewClient(obj.AlarmInfoPlate.IPAddr, 8131); nil != err {
			utils.DefaultSwitchLogger.Printf("create tcp client failed %v\n", err)
		} else {
			if cmd, err := cli.ScreenShowAndSayPrice(time.Now().Local().String(), "停车1小时27分钟", "收费17元"); nil != err {
				utils.DefaultSwitchLogger.Printf("show price on screen failed %v\n", err)
			} else {
				utils.DefaultSwitchLogger.Printf("show price on screen result %v\n", cmd.DataString())
			}

			if len(ret.VoiceContent) > 0 {
				if cmd, err := cli.TransmissionCmdSendKFVoice(ret.VoiceContent); nil != err {
					utils.DefaultSwitchLogger.Printf("----------- play voice on screen failed %v\n", err)
				} else {
					utils.DefaultSwitchLogger.Printf("----------- play voice on screen result %v\n", cmd.DataString())
				}
			}
		}
	}()
}
