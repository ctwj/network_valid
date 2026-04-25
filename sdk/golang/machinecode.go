package sdk

import (
	"crypto/md5"
	"encoding/hex"
	"net"
	"os"
)

// DefaultMachineCode 获取默认机器码
// 算法: MD5(MAC地址 + 主机名) 前 16 位
func DefaultMachineCode() string {
	mac := getMACAddress()
	hostname := getHostname()

	data := mac + hostname
	hash := md5.Sum([]byte(data))
	fullHash := hex.EncodeToString(hash[:])

	// 取前 16 位
	if len(fullHash) >= 16 {
		return fullHash[:16]
	}
	return fullHash
}

// getMACAddress 获取本机第一个非回环网络接口的 MAC 地址
func getMACAddress() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "unknown"
	}

	for _, iface := range interfaces {
		// 跳过回环接口
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		// 跳过未启用的接口
		if iface.Flags&net.FlagUp == 0 {
			continue
		}

		hwAddr := iface.HardwareAddr.String()
		if hwAddr != "" {
			return hwAddr
		}
	}

	return "unknown"
}

// getHostname 获取主机名
func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}
