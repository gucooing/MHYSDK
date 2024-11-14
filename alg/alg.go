package alg

import (
	"encoding/hex"
	"math/rand"
	"strconv"
)

func S2U32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}

func GetRandomByte(len int) []byte {
	ret := make([]byte, 0)
	for i := 0; i < len; i++ {
		r := uint8(rand.Intn(256))
		ret = append(ret, r)
	}
	return ret
}

func GetRandomByteHexStr(len int) string {
	return hex.EncodeToString(GetRandomByte(len))
}
