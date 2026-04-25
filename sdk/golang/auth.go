package sdk

const (
	endpointRegister = "/api/register"
	endpointLogin    = "/api/login"
	endpointHeartbeat = "/api/heartbeat"
	endpointUnbind   = "/api/unbind"
	endpointLogout   = "/api/logout"
)

// Register 用户注册
func (c *Client) Register(req RegisterRequest) error {
	params := map[string]string{
		"username": req.Username,
		"password": req.Password,
	}
	if req.Code != "" {
		params["code"] = req.Code
	}
	if req.Captcha != "" {
		params["captcha"] = req.Captcha
	}

	_, err := c.doMultipartRequest(endpointRegister, params)
	return err
}

// Login 用户登录
func (c *Client) Login(req LoginRequest) (*UserInfo, error) {
	params := map[string]string{
		"username": req.Username,
		"password": req.Password,
	}

	resp, err := c.doMultipartRequest(endpointLogin, params)
	if err != nil {
		return nil, err
	}

	return parseData[UserInfo](resp)
}

// Heartbeat 用户心跳
func (c *Client) Heartbeat(username string) error {
	params := map[string]string{
		"username": username,
	}

	_, err := c.doMultipartRequest(endpointHeartbeat, params)
	return err
}

// Unbind 用户解绑
func (c *Client) Unbind(username string) error {
	params := map[string]string{
		"username": username,
	}

	_, err := c.doMultipartRequest(endpointUnbind, params)
	return err
}

// Logout 用户下线
func (c *Client) Logout(username string) error {
	params := map[string]string{
		"username": username,
	}

	_, err := c.doMultipartRequest(endpointLogout, params)
	return err
}
