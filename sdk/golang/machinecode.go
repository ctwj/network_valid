package sdk

import (
	"crypto/md5"
	"encoding/hex"
	"net"
	"os"
	"sort"
	"strings"
)

// DefaultMachineCode 获取默认机器码
// 算法: MD5(所有有效MAC地址排序拼接 + 主机名 + 用户名)
// 确保固定性和唯一性
func DefaultMachineCode() string {
	macList := getAllMACAddresses()
	hostname := getHostname()
	username := getUsername()

	// 排序 MAC 地址列表确保固定性
	sort.Strings(macList)
	macCombined := strings.Join(macList, ",")

	data := macCombined + "|" + hostname + "|" + username
	hash := md5.Sum([]byte(data))
	fullHash := hex.EncodeToString(hash[:])

	// 取前 32 位作为机器码
	if len(fullHash) >= 32 {
		return fullHash[:32]
	}
	return fullHash
}

// getAllMACAddresses 获取所有有效网络接口的 MAC 地址
// 返回排序后的列表确保固定性
func getAllMACAddresses() []string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return []string{"unknown"}
	}

	var macList []string
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
		if hwAddr != "" && hwAddr != "00:00:00:00:00:00" {
			macList = append(macList, hwAddr)
		}
	}

	if len(macList) == 0 {
		return []string{"unknown"}
	}

	return macList
}

// getHostname 获取主机名
func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}

// getUsername 获取当前系统用户名
func getUsername() string {
	// 尝试从环境变量获取用户名
	username := os.Getenv("USER")
	if username == "" {
		username = os.Getenv("USERNAME") // Windows
	}
	if username == "" {
		username = os.Getenv("LOGNAME") // Unix
	}
	if username == "" {
		username = "unknown"
	}
	return username
}
