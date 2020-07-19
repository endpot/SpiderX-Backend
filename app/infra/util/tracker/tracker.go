package tracker

import (
	"encoding/hex"
)

// 将从 URL 获取到的字符串还原为 Hex 字符串
func RestoreToHexString(str string) string {
	s := hex.EncodeToString([]byte(str))
	return s
}

// 将 Hex 字符串再还原为 20-byte 字符串
// TODO: 优化代码
func RestoreToByteString(str string) string {
	byteSlice, err := hex.DecodeString(str)
	if err != nil {
		return str
	}

	return string(byteSlice)
}
