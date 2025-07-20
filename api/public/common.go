package public

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 生成32位MD5摘要
func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}
