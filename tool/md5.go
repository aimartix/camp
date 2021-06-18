package tool

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func GetMD5Encode(data string) string { //md5方法加密
	h := md5.New()
	h.Write([]byte(data))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

