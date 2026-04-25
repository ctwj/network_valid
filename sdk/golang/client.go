package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Client SDK 客户端
type Client struct {
	baseURL     string
	appKey      string
	secretKey   string
	version     string
	machineCode string
	httpClient  *http.Client
}

// NewClient 创建 SDK 客户端
func NewClient(cfg Config) (*Client, error) {
	// 验证必填参数
	if cfg.BaseURL == "" {
		return nil, fmt.Errorf("BaseURL is required")
	}
	if cfg.AppKey == "" {
		return nil, fmt.Errorf("AppKey is required")
	}
	if cfg.SecretKey == "" {
		return nil, fmt.Errorf("SecretKey is required")
	}
	if cfg.Version == "" {
		return nil, fmt.Errorf("Version is required")
	}

	// 处理机器码
	machineCode := cfg.MachineCode
	if machineCode == "" {
		machineCode = DefaultMachineCode()
	}

	// 处理 HTTP Client
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	// 规范化 BaseURL
	baseURL := strings.TrimSuffix(cfg.BaseURL, "/")

	return &Client{
		baseURL:     baseURL,
		appKey:      cfg.AppKey,
		secretKey:   cfg.SecretKey,
		version:     cfg.Version,
		machineCode: machineCode,
		httpClient:  httpClient,
	}, nil
}

// buildCommonParams 构建公共参数
func (c *Client) buildCommonParams(timestamp int64) url.Values {
	params := url.Values{}
	params.Set("appkey", c.appKey)
	params.Set("timestamp", strconv.FormatInt(timestamp, 10))
	params.Set("version", c.version)
	params.Set("mac", c.machineCode)

	// 生成签名
	sign := GenerateSign(c.appKey, c.secretKey, c.version, strconv.FormatInt(timestamp, 10), c.machineCode)
	params.Set("sign", sign)

	return params
}

// doFormRequest 发送 application/x-www-form-urlencoded 请求
func (c *Client) doFormRequest(endpoint string, extraParams url.Values) (*Response, error) {
	timestamp := time.Now().Unix()
	params := c.buildCommonParams(timestamp)

	// 合并额外参数
	if extraParams != nil {
		for key, values := range extraParams {
			for _, value := range values {
				params.Add(key, value)
			}
		}
	}

	reqURL := c.baseURL + endpoint
	req, err := http.NewRequest("POST", reqURL, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return c.doRequest(req)
}

// doMultipartRequest 发送 multipart/form-data 请求
func (c *Client) doMultipartRequest(endpoint string, extraParams map[string]string) (*Response, error) {
	timestamp := time.Now().Unix()
	commonParams := c.buildCommonParams(timestamp)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// 写入公共参数
	for key, values := range commonParams {
		for _, value := range values {
			_ = writer.WriteField(key, value)
		}
	}

	// 写入额外参数
	if extraParams != nil {
		for key, value := range extraParams {
			_ = writer.WriteField(key, value)
		}
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("close multipart writer failed: %w", err)
	}

	reqURL := c.baseURL + endpoint
	req, err := http.NewRequest("POST", reqURL, &body)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return c.doRequest(req)
}

// doRequest 执行 HTTP 请求并解析响应
func (c *Client) doRequest(req *http.Request) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %w", err)
	}

	var apiResp Response
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("parse response failed: %w", err)
	}

	// 检查业务错误
	if apiResp.Errno != 0 {
		return nil, &APIError{
			Errno:     apiResp.Errno,
			Errmsg:    apiResp.Errmsg,
			UID:       apiResp.UID,
			Timestamp: apiResp.Timestamp,
		}
	}

	return &apiResp, nil
}

// parseData 解析响应数据
func parseData[T any](resp *Response) (*T, error) {
	if resp == nil || len(resp.Data) == 0 {
		return nil, nil
	}

	var data T
	if err := json.Unmarshal(resp.Data, &data); err != nil {
		return nil, fmt.Errorf("parse data failed: %w", err)
	}

	return &data, nil
}
