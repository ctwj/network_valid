package sdk

import "strconv"

const (
	endpointRecharge = "/api/recharge"
	endpointDeduct   = "/api/deduct"
	endpointBan      = "/api/ban"
	endpointIsOnline = "/api/online"
	endpointRecover  = "/api/recover"
)

// Recharge 用户充值
func (c *Client) Recharge(req RechargeRequest) error {
	params := map[string]string{
		"username": req.Username,
		"amount":   strconv.Itoa(req.Amount),
	}

	_, err := c.doMultipartRequest(endpointRecharge, params)
	return err
}

// Deduct 账号扣点
func (c *Client) Deduct(req DeductRequest) error {
	params := map[string]string{
		"username": req.Username,
		"amount":   strconv.Itoa(req.Amount),
	}

	_, err := c.doMultipartRequest(endpointDeduct, params)
	return err
}

// Ban 账号拉黑
func (c *Client) Ban(req BanRequest) error {
	params := map[string]string{
		"username": req.Username,
	}
	if req.Reason != "" {
		params["reason"] = req.Reason
	}

	_, err := c.doMultipartRequest(endpointBan, params)
	return err
}

// IsOnline 查询用户在线状态
func (c *Client) IsOnline(username string) (*OnlineStatus, error) {
	params := map[string]string{
		"username": username,
	}

	resp, err := c.doMultipartRequest(endpointIsOnline, params)
	if err != nil {
		return nil, err
	}

	return parseData[OnlineStatus](resp)
}

// Recover 找回账号
func (c *Client) Recover(req RecoverRequest) error {
	params := map[string]string{
		"email":    req.Email,
		"code":     req.Code,
		"captcha":  req.Captcha,
	}

	_, err := c.doMultipartRequest(endpointRecover, params)
	return err
}
