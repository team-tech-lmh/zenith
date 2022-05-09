package tcpsdk

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"zenith/zenithsdk/utils"

	"github.com/sigurn/crc16"
)

type TransmissionCmdBody struct {
	Cmd     string `json:"cmd"`
	Subcmd  string `json:"subcmd"`
	Datalen int    `json:"datalen"`
	Comm    string `json:"comm"`
	Data    string `json:"data"`
}

func NewKFVoiceCmd(voiceContent string) (*TransmissionCmdBody, error) {
	cmdStruct := TransmissionCmdBody{
		Cmd:    "ttransmission",
		Subcmd: "send",
		Comm:   "rs485-1",
	}
	if buf, err := utils.Utf8ToGbk([]byte(voiceContent)); nil != err {
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

		return &cmdStruct, nil
	}
}
