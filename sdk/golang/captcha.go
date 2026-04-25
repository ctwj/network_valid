package sdk

import (
	"net/url"
)

const (
	endpointCaptcha          = "/api/captcha"
	endpointRegisterCode     = "/api/code/register"
	endpointRecoverCode      = "/api/code/recover"
)

// GetCaptcha 获取图形验证码
func (c *Client) GetCaptcha() ([]byte, error) {
	resp, err := c.doFormRequest(endpointCaptcha, nil)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// SendRegisterCode 发送注册验证码邮件
func (c *Client) SendRegisterCode(email string) error {
	params := url.Values{}
	params.Set("email", email)

	_, err := c.doFormRequest(endpointRegisterCode, params)
	return err
}

// SendRecoverCode 发送找回账号验证码邮件
func (c *Client) SendRecoverCode(email string) error {
	params := url.Values{}
	params.Set("email", email)

	_, err := c.doFormRequest(endpointRecoverCode, params)
	return err
}
