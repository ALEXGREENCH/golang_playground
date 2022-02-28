package main

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	//"math/rand"// Что-то не очень рандомно
	//"bytes"
	//"math/bits"
	//"strconv"
)

const (
	flapStartMarker byte = '\x2a'
)

const (
	flapFrameSignOn    byte = '\x01'
	flapFrameData      byte = '\x02'
	flapFrameError     byte = '\x03'
	flapFrameSignOff   byte = '\x04'
	flapFrameKeepAlive byte = '\x05'
)

const (
	flapMulticonnOldClient    byte = '\x00'
	flapMulticonnRecentClient byte = '\x01'
	flapMulticonnSingleClient byte = '\x02'
)

var (
	flap_1_1_value = []byte{'\x00', '\x00', '\x00', '\x01'}
)

func main() {

	var fullWellcomeFlap = getSendData(flapFrameSignOn, flap_1_1_value)

	fmt.Printf("%s", hex.Dump(fullWellcomeFlap))
}

func genRandomSeq() []byte {
	token := make([]byte, 2)
	rand.Read(token)
	return token
}

func getSendData(chanel byte, data []byte) []byte {
	tokenAndChanel := []byte{flapStartMarker, flapFrameSignOn}

	flapHeader := append(tokenAndChanel, genRandomSeq()...)
	fullFlap := append(flapHeader, lenghtAddDataBytes(data)...)

	return fullFlap
}

func createTVL(tag []byte, data []byte) []byte {
	tvl := append(tag, lenghtAddDataBytes(data)...)
	return tvl
}

func lenghtAddDataBytes(data []byte) []byte {

	// сдвиг на 2 байта, толку нету, как пример :)
	s := make([]byte, len(data)+2)
	for k, v := range data {
		s[k+2] = v
	}

	l := calDataLenght(data)
	s[0] = l[0]
	s[1] = l[1]

	return s
}

// по спецификации длинна не может привышать 2 байта
func calDataLenght(data []byte) []byte {
	i := int8(len(data))
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(i))
	return b
}
