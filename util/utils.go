package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func BytesToString(bytes []byte) string {
	return fmt.Sprintf("%x", bytes)
}
