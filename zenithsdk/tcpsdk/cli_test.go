package tcpsdk

import (
	"log"
	"testing"
)

func Test_Voice(t *testing.T) {
	if cli, err := NewClient("192.168.2.233", 8131); nil != err {
		log.Printf("create tcp client failed %v\n", err)
	} else {
		if cmd, err := cli.TransmissionCmdSend("欢迎光临"); nil != err {
			log.Printf("show price on screen failed %v\n", err)
		} else {
			log.Printf("show price on screen result %v\n", cmd.DataString())
		}
	}
}
