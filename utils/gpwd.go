package utils

import (
	"encoding/hex"
	"crypto/md5"
)

//Md5encode 密码md5加密
func Md5encode (key string) string {
	hash := md5.New()
	hash.Write([]byte(key))
	return hex.EncodeToString(hash.Sum(nil))
}

//Md5Compare 密码校验，比对结果
func Md5Compare (str string,key string) bool {
	if Md5encode(str) == key {
		return true
	}
	return false
}