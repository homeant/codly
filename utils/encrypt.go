package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// EncryptPassword 使用 bcrypt 对密码进行加密
func EncryptPassword(password string, key string) string {
	// 创建一个新的 HMAC 使用 SHA-256 哈希函数
	h := hmac.New(sha256.New, []byte(key))
	// 对密码进行哈希
	// 写入数据
	h.Write([]byte(password))
	// 返回生成的 HMAC 值（十六进制表示）
	return hex.EncodeToString(h.Sum(nil))
}
