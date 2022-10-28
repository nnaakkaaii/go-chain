package casttype

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)

	if err := binary.Write(buff, binary.BigEndian, num); err != nil {
		log.Panicln(err)
	}

	return buff.Bytes()
}
