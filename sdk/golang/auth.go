package sdk

import "fmt"

const (
	endpointAPI      = "/api/index/"
	actionRegister   = "user.register"
	actionLogin      = "user.login"
	actionHeartbeat  = "user.heart"
	actionUnbind     = "user.bind"
	actionLogout     = "user.logout"
)

// Register 用户注册
func (c *Client) Register(req RegisterRequest) error {
	params := map[string]string{
		"action": actionRegister,
		"user":   req.Username,
		"pwd":    req.Password,
	}
	if req.Code != "" {
		params["code"] = req.Code
	}
	if req.Captcha != "" {
		params["captcha"] = req.Captcha
	}

	_, err := c.doMultipartRequest(endpointAPI, params)
	return err
}

// Login 用户登录
func (c *Client) Login(req LoginRequest) (*UserInfo, error) {
	params := map[string]string{
		"action": actionLogin,
		"user":   req.Username,
		"pwd":    req.Password,
	}

	resp, err := c.doMultipartRequest(endpointAPI, params)
	if err != nil {
		return nil, err
	}

	userInfo, err := parseData[UserInfo](resp)
	if err != nil {
		return nil, err
	}

	// 保存 client token 用于后续心跳
	if userInfo != nil && userInfo.Client != "" {
		c.clientToken = userInfo.Client
	}

	return userInfo, nil
}

// Heartbeat 用户心跳（使用登录返回的 client token）
func (c *Client) Heartbeat() error {
	if c.clientToken == "" {
		return fmt.Errorf("client token is empty, please login first")
	}

	params := map[string]string{
		"action": actionHeartbeat,
		"client": c.clientToken,
	}

	_, err := c.doMultipartRequest(endpointAPI, params)
	return err
}

// HeartbeatWithClient 使用指定的 client token 发送心跳
func (c *Client) HeartbeatWithClient(clientToken string) error {
	if clientToken == "" {
		return fmt.Errorf("client token is required")
	}

	params := map[string]string{
		"action": actionHeartbeat,
		"client": clientToken,
	}

	_, err := c.doMultipartRequest(endpointAPI, params)
	return err
}

// Unbind 用户解绑
func (c *Client) Unbind(username string) error {
	params := map[string]string{
		"action":   actionUnbind,
		"username": username,
	}

	_, err := c.doMultipartRequest(endpointAPI, params)
	return err
}

// Logout 用户下线
func (c *Client) Logout() error {
	if c.clientToken == "" {
		return fmt.Errorf("client token is empty, please login first")
	}

	params := map[string]string{
		"action": actionLogout,
		"client": c.clientToken,
		"type":   "0", // 0: 当前客户端下线, 1: 全部下线
	}

	_, err := c.doMultipartRequest(endpointAPI, params)
	return err
}

// LogoutAll 下线所有客户端
func (c *Client) LogoutAll() error {
	if c.clientToken == "" {
		return fmt.Errorf("client token is empty, please login first")
	}

	params := map[string]string{
		"action": actionLogout,
		"client": c.clientToken,
		"type":   "1", // 0: 当前客户端下线, 1: 全部下线
	}

	_, err := c.doMultipartRequest(endpointAPI, params)
	return err
}
