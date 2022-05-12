package tcpsdk

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"

	"github.com/xizhukarsa/zenith/utils"

	"github.com/sigurn/crc16"
)

func Test_Voice(t *testing.T) {
	if cli, err := NewClient("192.168.2.233", 8131); nil != err {
		log.Printf("create tcp client failed %v\n", err)
	} else {
		if cmd, err := cli.TransmissionCmdSendKFVoice("月租车剩余"); nil != err {
			log.Printf("show price on screen failed %v\n", err)
		} else {
			log.Printf("show price on screen result %v\n", cmd.DataString())
		}
	}
}

func Test_CRC16(t *testing.T) {
	hexstr := "0064ffff300401313233"
	buf, err := hex.DecodeString(hexstr)
	if nil != err {
		t.Error(err)
		return
	}

	table := crc16.MakeTable(crc16.CRC16_MODBUS)
	sum := crc16.Checksum(buf, table)
	bu1f, _ := utils.Uint16(sum).ToBytes()
	fmt.Printf("%v\n", bu1f)
}

func Test_Close(t *testing.T) {
	if cli, err := NewClient("192.168.2.237", 8131); nil != err {
		log.Printf("create tcp client failed %v\n", err)
	} else {
		if cmd, err := cli.OpenBarrier(); nil != err {
			log.Printf("show price on screen failed %v\n", err)
		} else {
			log.Printf("show price on screen result %v\n", cmd.DataString())
		}
		if cmd, err := cli.CloseBarrier(); nil != err {
			log.Printf("show price on screen failed %v\n", err)
		} else {
			log.Printf("show price on screen result %v\n", cmd.DataString())
		}
	}
}
