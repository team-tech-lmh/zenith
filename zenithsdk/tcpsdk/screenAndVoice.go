package tcpsdk

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"zenith/utils"

	"github.com/sigurn/crc16"
)

type ScreenShowItem struct {
	Show_mode    int    `json:"show_mode"`
	Show_content string `json:"show_content"`
}

type CustomVoice struct {
	Voice_volume int    `json:"voice_volume"`
	Play_content string `json:"play_content"`
}

type VoiceConfig struct {
	Voice_mode              int `json:"voice_mode"`
	Voice_welcom            int `json:"voice_welcom"`
	Voice_tag               int `json:"voice_tag"`
	Temporary_voice_welcome int `json:"temporary_voice_welcome"`
	Temporary_voice_tag     int `json:"temporary_voice_tag"`
}

type ScreenShowAndSayPriceCmdBody struct {
	Screen_ctrl_pro_type int              `json:"screen_ctrl_pro_type"`
	Use_serial_port      int              `json:"use_serial_port"`
	Screen_isopen        int              `json:"screen_isopen"`
	Free_cfg             []ScreenShowItem `json:"free_cfg"`
	Busy_cfg             []ScreenShowItem `json:"busy_cfg"`
	Voice_cfg            VoiceConfig      `json:"voice_cfg"`
	Show_direction       int              `json:"show_direction"`
}

func (cli *Client) ScreenShowAndSayPrice(stayDuStr, priceStr string) (*Cmd, error) {
	cmdBody := ScreenShowAndSayPriceCmdBody{
		Screen_ctrl_pro_type: 3,
		Use_serial_port:      0,
		Screen_isopen:        2,
		Free_cfg: []ScreenShowItem{
			{
				Show_mode: 2, //当前时间
			},
			{
				Show_mode:    1,
				Show_content: base64Str("小镇希望"),
			},
			{
				Show_mode:    1,
				Show_content: base64Str("联每户"),
			},
			{
				Show_mode:    1,
				Show_content: base64Str("欢迎回家"),
			},
		},
		Busy_cfg: []ScreenShowItem{
			{
				Show_mode: 8, //车牌号
			},
			{
				Show_mode:    1,
				Show_content: base64Str(priceStr),
			},
			{
				Show_mode:    1,
				Show_content: base64Str(stayDuStr),
			},
			{
				Show_mode:    1,
				Show_content: base64Str("一路顺风"),
			},
		},
		Voice_cfg: VoiceConfig{
			Voice_mode:              1,
			Voice_welcom:            1,
			Voice_tag:               4,
			Temporary_voice_welcome: 1,
			Temporary_voice_tag:     7,
		},
		Show_direction: 0,
	}
	cmdMap := map[string]interface{}{
		"cmd":  "set_led_ctrl_cfg",
		"body": cmdBody,
	}
	cmdBuf, err := json.Marshal(cmdMap)
	if nil != err {
		return nil, err
	}
	cmd := NewDataCmd(cmdBuf, 0, CmdTypeData)
	if err := cli.SendCmd(cmd); nil != err {
		return nil, err
	}

	if c, err := cli.ReceiveCmd(); nil != err {
		return nil, err
	} else {
		return c, nil
	}
}

func (cli *Client) TransmissionCmdSend(data string) (*Cmd, error) {
	type TransmissionCmdBody struct {
		Cmd     string `json:"cmd"`
		Subcmd  string `json:"subcmd"`
		Datalen int    `json:"datalen"`
		Comm    string `json:"comm"`
		Data    string `json:"data"`
	}

	cmdStruct := TransmissionCmdBody{
		Cmd:    "ttransmission",
		Subcmd: "send",
		Comm:   "rs485-1",
	}
	if buf, err := utils.Utf8ToGbk([]byte(data)); nil != err {
		return nil, err
	} else {
		buf = append([]byte{0x01}, buf...)

		lbuf, err := utils.Uint8(len(buf)).ToBytes()
		if nil != err {
			return nil, err
		}

		hexCmdHead := "0064FFFF30"
		hexBuf, err := hex.DecodeString(hexCmdHead)
		if nil != err {
			return nil, err
		}
		hexBuf = append(hexBuf, lbuf...)
		buf = append(hexBuf, buf...)

		fmt.Println(hex.EncodeToString(buf))

		table := crc16.MakeTable(crc16.CRC16_MODBUS)
		sum := crc16.Checksum(buf, table)
		sumBuf, err := utils.Uint16(sum).ToBytes()
		if nil != err {
			return nil, err
		}
		buf = append(buf, sumBuf...)

		cmdStruct.Datalen = len(buf)
		d := base64.StdEncoding.EncodeToString(buf)
		cmdStruct.Data = d
	}

	cmdBuf, err := json.Marshal(cmdStruct)
	if nil != err {
		return nil, err
	}
	cmd := NewDataCmd(cmdBuf, 0, CmdTypeData)
	if err := cli.SendCmd(cmd); nil != err {
		return nil, err
	}
	if c, err := cli.ReceiveCmd(); nil != err {
		return nil, err
	} else {
		return c, nil
	}
}

func base64Str(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
