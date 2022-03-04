package zenithsdk

import (
	"encoding/binary"
	"encoding/json"
	"log"
)

const (
	cmdHederLen = 8
)

type CmdType uint8

const (
	CmdTypeData      CmdType = 0x00
	CmdTypeHeartBeat CmdType = 0x01
)

type CmdHeader struct {
	Lead1        uint8
	Lead2        uint8
	PackageType  CmdType
	PackageIndex uint8
	PackageLen   uint32
}

func CmdHeaderFromBuf(buf8 []byte) CmdHeader {
	if len(buf8) != 8 {
		panic("cmd header construct failed with wrong buf length")
	}
	cmd := CmdHeader{}
	cmd.Lead1 = uint8(buf8[0])
	cmd.Lead2 = uint8(buf8[1])
	cmd.PackageType = CmdType(buf8[2])
	cmd.PackageIndex = uint8(buf8[3])
	cmd.PackageLen = binary.BigEndian.Uint32(buf8[4:8])

	return cmd
}

type Cmd struct {
	CmdHeader
	Data []byte
}

func NewDataCmd(data []byte, idx uint8, pt CmdType) Cmd {
	ret := Cmd{
		CmdHeader: CmdHeader{
			Lead1:        uint8('V'),
			Lead2:        uint8('Z'),
			PackageType:  pt,
			PackageIndex: idx,
			PackageLen:   uint32(len(data)),
		},
		Data: data,
	}
	return ret
}

func (cmd Cmd) ToBytes() []byte {
	ret := []byte{}
	ret = append(ret, byte(cmd.Lead1))
	ret = append(ret, byte(cmd.Lead2))
	ret = append(ret, byte(cmd.PackageType))
	ret = append(ret, byte(cmd.PackageIndex))

	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, cmd.PackageLen)
	ret = append(ret, buf...)

	ret = append(ret, cmd.Data...)
	return ret
}

func (cmd Cmd) DataMap() (map[string]interface{}, error) {
	var ret map[string]interface{}
	if err := json.Unmarshal(cmd.Data, &ret); nil != err {
		log.Printf("cmd data to json map failed %v\n", err)
		return nil, err
	}
	return ret, nil
}

func (cmd Cmd) DataString() string {
	return string(cmd.Data)
}
