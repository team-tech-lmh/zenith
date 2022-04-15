package utils

import (
	"bytes"
	"encoding/binary"
)

type ToBytes interface {
	ToBytes() []byte
}

type Uint16 uint16

func (i Uint16) ToBytes() ([]byte, error) {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.LittleEndian, uint16(i))
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

type Uint8 uint8

func (i Uint8) ToBytes() ([]byte, error) {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, uint8(i))
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}
