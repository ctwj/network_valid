package sdk_test

import (
	"fmt"
	"time"

	"github.com/kerwin/network_valid/sdk/golang"
)

func ExampleNewClient() {
	// 创建 SDK 客户端（自动获取机器码）
	client, err := sdk.NewClient(sdk.Config{
		BaseURL:   "https://api.example.com",
		AppKey:    "your_appkey",
		SecretKey: "your_secretkey",
		Version:   "1.0.0",
		// MachineCode 不传，SDK 自动获取
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Client created successfully\n")
	_ = client
}

func ExampleNewClient_withCustomMachineCode() {
	// 创建 SDK 客户端（自定义机器码）
	client, err := sdk.NewClient(sdk.Config{
		BaseURL:     "https://api.example.com",
		AppKey:      "your_appkey",
		SecretKey:   "your_secretkey",
		Version:     "1.0.0",
		MachineCode: "my_custom_machine_code",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Client created with custom machine code\n")
	_ = client
}

func ExampleClient_Login() {
	client, _ := sdk.NewClient(sdk.Config{
		BaseURL:   "https://api.example.com",
		AppKey:    "your_appkey",
		SecretKey: "your_secretkey",
		Version:   "1.0.0",
	})

	// 用户登录
	userInfo, err := client.Login(sdk.LoginRequest{
		Username: "user@example.com",
		Password: "password123",
	})
	if err != nil {
		// 处理错误
		if apiErr, ok := err.(*sdk.APIError); ok {
			fmt.Printf("登录失败: %s (错误码: %d)\n", apiErr.Errmsg, apiErr.Errno)
		}
		return
	}

	fmt.Printf("登录成功: %s\n", userInfo.Username)
}

func ExampleClient_Register() {
	client, _ := sdk.NewClient(sdk.Config{
		BaseURL:   "https://api.example.com",
		AppKey:    "your_appkey",
		SecretKey: "your_secretkey",
		Version:   "1.0.0",
	})

	// 用户注册
	err := client.Register(sdk.RegisterRequest{
		Username: "newuser@example.com",
		Password: "password123",
	})
	if err != nil {
		fmt.Printf("注册失败: %v\n", err)
		return
	}

	fmt.Printf("注册成功\n")
}

func ExampleClient_Heartbeat() {
	client, _ := sdk.NewClient(sdk.Config{
		BaseURL:   "https://api.example.com",
		AppKey:    "your_appkey",
		SecretKey: "your_secretkey",
		Version:   "1.0.0",
	})

	// 先登录获取 client token
	_, err := client.Login(sdk.LoginRequest{
		Username: "user@example.com",
		Password: "password123",
	})
	if err != nil {
		fmt.Printf("登录失败: %v\n", err)
		return
	}

	// 定期发送心跳（建议每 30 秒）
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if err := client.Heartbeat(); err != nil {
			fmt.Printf("心跳失败: %v\n", err)
			// 心跳失败可能需要重新登录
			break
		}
	}
}

func ExampleClient_GetPermissions() {
	client, _ := sdk.NewClient(sdk.Config{
		BaseURL:   "https://api.example.com",
		AppKey:    "your_appkey",
		SecretKey: "your_secretkey",
		Version:   "1.0.0",
	})

	// 获取用户权限列表
	permissions, err := client.GetPermissions()
	if err != nil {
		fmt.Printf("获取权限失败: %v\n", err)
		return
	}

	for _, perm := range permissions {
		fmt.Printf("权限: %s (%s)\n", perm.Name, perm.Path)
	}
}

func ExampleDefaultMachineCode() {
	// 获取默认机器码
	machineCode := sdk.DefaultMachineCode()
	fmt.Printf("机器码: %s\n", machineCode)
}
