package sdk

import (
	"crypto/md5"
	"encoding/hex"
)

// GenerateSign 生成 MD5 签名
// 签名算法: MD5(appkey + secretkey + version + timestamp + mac)
func GenerateSign(appkey, secretkey, version, timestamp, mac string) string {
	signStr := appkey + secretkey + version + timestamp + mac
	hash := md5.Sum([]byte(signStr))
	return hex.EncodeToString(hash[:])
}

// GenerateSignBytes 生成 MD5 签名（字节版本）
func GenerateSignBytes(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
