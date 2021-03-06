package tcpsdk

import (
	"encoding/hex"
	"testing"

	"github.com/team-tech-lmh/zenith/utils"

	"github.com/sigurn/crc16"
)

func Test_Voice(t *testing.T) {
	if cli, err := NewClient("192.168.2.233", 8131); nil != err {
		utils.DefaultSwitchLogger.Printf("create tcp client failed %v\n", err)
	} else {
		if cmd, err := cli.TransmissionCmdSendKFVoice("月租车剩余"); nil != err {
			utils.DefaultSwitchLogger.Printf("show price on screen failed %v\n", err)
		} else {
			utils.DefaultSwitchLogger.Printf("show price on screen result %v\n", cmd.DataString())
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
	utils.DefaultSwitchLogger.Printf("%v\n", bu1f)
}

func Test_Close(t *testing.T) {
	if cli, err := NewClient("192.168.2.237", 8131); nil != err {
		utils.DefaultSwitchLogger.Printf("create tcp client failed %v\n", err)
	} else {
		if cmd, err := cli.OpenBarrier(); nil != err {
			utils.DefaultSwitchLogger.Printf("show price on screen failed %v\n", err)
		} else {
			utils.DefaultSwitchLogger.Printf("show price on screen result %v\n", cmd.DataString())
		}
		if cmd, err := cli.CloseBarrier(); nil != err {
			utils.DefaultSwitchLogger.Printf("show price on screen failed %v\n", err)
		} else {
			utils.DefaultSwitchLogger.Printf("show price on screen result %v\n", cmd.DataString())
		}
	}
}
