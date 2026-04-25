package sdk

import (
	"encoding/json"
	"time"
)

const (
	endpointTimestamp      = "/api/timestamp"
	endpointPermissions    = "/api/permissions"
	endpointSoftwareInfo   = "/api/software/info"
	endpointMemberTags     = "/api/member/tags"
	endpointRemoteVariables = "/api/remote/variables"
)

// GetTimestamp 获取服务器时间戳
func (c *Client) GetTimestamp() (int64, error) {
	resp, err := c.doFormRequest(endpointTimestamp, nil)
	if err != nil {
		return 0, err
	}

	var timestamp int64
	if err := json.Unmarshal(resp.Data, &timestamp); err != nil {
		return 0, err
	}

	return timestamp, nil
}

// GetPermissions 获取用户权限列表
func (c *Client) GetPermissions() ([]Permission, error) {
	resp, err := c.doMultipartRequest(endpointPermissions, nil)
	if err != nil {
		return nil, err
	}

	data, err := parseData[[]Permission](resp)
	if err != nil {
		return nil, err
	}
	return *data, nil
}

// GetSoftwareInfo 获取软件信息
func (c *Client) GetSoftwareInfo() (*SoftwareInfo, error) {
	resp, err := c.doMultipartRequest(endpointSoftwareInfo, nil)
	if err != nil {
		return nil, err
	}

	return parseData[SoftwareInfo](resp)
}

// GetMemberTags 获取会员标签列表
func (c *Client) GetMemberTags() ([]MemberTag, error) {
	resp, err := c.doMultipartRequest(endpointMemberTags, nil)
	if err != nil {
		return nil, err
	}

	data, err := parseData[[]MemberTag](resp)
	if err != nil {
		return nil, err
	}
	return *data, nil
}

// GetRemoteVariables 获取远程变量
func (c *Client) GetRemoteVariables() (map[string]string, error) {
	resp, err := c.doMultipartRequest(endpointRemoteVariables, nil)
	if err != nil {
		return nil, err
	}

	// 远程变量可能是键值对数组或对象
	var variables map[string]string
	if err := json.Unmarshal(resp.Data, &variables); err != nil {
		// 尝试解析为数组
		var arr []RemoteVariable
		if err := json.Unmarshal(resp.Data, &arr); err != nil {
			return nil, err
		}
		variables = make(map[string]string)
		for _, v := range arr {
			variables[v.Key] = v.Value
		}
	}

	return variables, nil
}

// GetRemoteVariable 获取单个远程变量
func (c *Client) GetRemoteVariable(key string) (string, error) {
	variables, err := c.GetRemoteVariables()
	if err != nil {
		return "", err
	}

	if value, ok := variables[key]; ok {
		return value, nil
	}

	return "", nil
}

// SyncTimestamp 同步服务器时间戳
// 返回服务器与本地时间的偏差（秒）
func (c *Client) SyncTimestamp() (int64, error) {
	localBefore := time.Now().Unix()
	serverTS, err := c.GetTimestamp()
	if err != nil {
		return 0, err
	}
	localAfter := time.Now().Unix()

	// 计算网络延迟
	latency := (localAfter - localBefore) / 2
	offset := serverTS - (localBefore + latency)

	return offset, nil
}

func getCurrentTimestamp() int64 {
	return time.Now().Unix()
}
