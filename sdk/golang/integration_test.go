package sdk_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/kerwin/network_valid/sdk/golang"
)

// 测试配置 - 请确保服务器上已创建对应版本号
var testConfig = sdk.Config{
	BaseURL:   "https://app.l9.lc",
	AppKey:    "DoBHMi5mTLi8r4wg18MtmqPmZeK3eGuq",
	SecretKey: "UcS1CM6OnXUEbsbAugyZgxsBo3X1iJgV",
	Version:   "1.00", // 版本号格式为 X.XX
}

// 测试账号（单码模式下使用激活码作为用户名，密码任意）
var testUsername = "798605218247" // 激活码
var testPassword = "any"          // 单码模式下密码任意

func TestMachineCode(t *testing.T) {
	machineCode := sdk.DefaultMachineCode()
	t.Logf("机器码: %s (长度: %d)", machineCode, len(machineCode))

	if len(machineCode) != 32 {
		t.Errorf("机器码长度应为 32，实际为 %d", len(machineCode))
	}

	// 验证多次调用结果一致（固定性）
	for i := 0; i < 3; i++ {
		code := sdk.DefaultMachineCode()
		if code != machineCode {
			t.Errorf("机器码不固定: 第 %d 次调用结果不同", i+1)
		}
	}
	t.Log("机器码固定性验证通过")
}

func TestClientCreation(t *testing.T) {
	client, err := sdk.NewClient(testConfig)
	if err != nil {
		t.Fatalf("创建客户端失败: %v", err)
	}
	t.Logf("客户端创建成功，机器码: %s", client.MachineCode())
}

func TestLogin(t *testing.T) {
	client, err := sdk.NewClient(testConfig)
	if err != nil {
		t.Fatalf("创建客户端失败: %v", err)
	}

	// 使用测试账号登录
	userInfo, err := client.Login(sdk.LoginRequest{
		Username: testUsername,
		Password: testPassword,
	})
	if err != nil {
		if apiErr, ok := err.(*sdk.APIError); ok {
			t.Logf("登录失败: %s (错误码: %d)", apiErr.Errmsg, apiErr.Errno)
			// 如果是账号不存在，尝试注册
			if apiErr.Errno == 400 {
				t.Log("尝试注册新账号...")
				regErr := client.Register(sdk.RegisterRequest{
					Username: testUsername,
					Password: testPassword,
				})
				if regErr != nil {
					t.Logf("注册失败: %v", regErr)
					t.Skip("无法注册测试账号")
				}
				t.Log("注册成功，重新登录...")
				userInfo, err = client.Login(sdk.LoginRequest{
					Username: testUsername,
					Password: testPassword,
				})
				if err != nil {
					t.Fatalf("登录失败: %v", err)
				}
			} else {
				t.Fatalf("登录失败: %v", err)
			}
		} else {
			t.Fatalf("登录失败: %v", err)
		}
	}

	if userInfo == nil {
		t.Fatal("用户信息为空")
	}

	t.Logf("登录成功!")
	t.Logf("  用户名: %s", userInfo.Username)
	t.Logf("  昵称: %s", userInfo.Nickname)
	t.Logf("  Client Token: %s", userInfo.Client)
	t.Logf("  到期时间: %s", userInfo.Endtime)
	t.Logf("  剩余天数: %d", userInfo.RealCdays)
	t.Logf("  剩余点数: %d", userInfo.CountPoints)
	t.Logf("  标签: %s", userInfo.Tag)

	if userInfo.Client == "" {
		t.Error("Client Token 为空，心跳将无法工作")
	}
}

func TestHeartbeat(t *testing.T) {
	client, err := sdk.NewClient(testConfig)
	if err != nil {
		t.Fatalf("创建客户端失败: %v", err)
	}

	// 登录
	userInfo, err := client.Login(sdk.LoginRequest{
		Username: testUsername,
		Password: testPassword,
	})
	if err != nil {
		t.Fatalf("登录失败: %v", err)
	}
	t.Logf("登录成功，Client: %s", userInfo.Client)

	// 发送心跳
	err = client.Heartbeat()
	if err != nil {
		t.Fatalf("心跳失败: %v", err)
	}
	t.Log("心跳成功!")

	// 连续发送多次心跳
	for i := 0; i < 3; i++ {
		time.Sleep(2 * time.Second)
		err = client.Heartbeat()
		if err != nil {
			t.Errorf("第 %d 次心跳失败: %v", i+1, err)
			break
		}
		t.Logf("第 %d 次心跳成功", i+1)
	}
}

func TestLogout(t *testing.T) {
	client, err := sdk.NewClient(testConfig)
	if err != nil {
		t.Fatalf("创建客户端失败: %v", err)
	}

	// 登录
	_, err = client.Login(sdk.LoginRequest{
		Username: testUsername,
		Password: testPassword,
	})
	if err != nil {
		t.Fatalf("登录失败: %v", err)
	}
	t.Log("登录成功")

	// 发送一次心跳确认在线
	err = client.Heartbeat()
	if err != nil {
		t.Fatalf("心跳失败: %v", err)
	}
	t.Log("心跳成功")

	// 下线
	err = client.Logout()
	if err != nil {
		t.Logf("下线失败: %v", err)
	} else {
		t.Log("下线成功")
	}

	// 下线后心跳应该失败
	time.Sleep(1 * time.Second)
	err = client.Heartbeat()
	if err != nil {
		t.Logf("下线后心跳失败（预期行为）: %v", err)
	} else {
		t.Error("下线后心跳仍然成功，不符合预期")
	}
}

func TestFullWorkflow(t *testing.T) {
	client, err := sdk.NewClient(testConfig)
	if err != nil {
		t.Fatalf("创建客户端失败: %v", err)
	}

	t.Log("=== 完整工作流测试 ===")

	// 1. 登录
	t.Log("1. 登录中...")
	userInfo, err := client.Login(sdk.LoginRequest{
		Username: testUsername,
		Password: testPassword,
	})
	if err != nil {
		t.Fatalf("登录失败: %v", err)
	}
	t.Logf("   登录成功! Client: %s", userInfo.Client)

	// 2. 模拟在线状态（发送 3 次心跳）
	t.Log("2. 保持在线状态...")
	for i := 0; i < 3; i++ {
		time.Sleep(2 * time.Second)
		if err = client.Heartbeat(); err != nil {
			t.Fatalf("心跳失败: %v", err)
		}
		t.Logf("   心跳 %d/3 成功", i+1)
	}

	// 3. 正常下线
	t.Log("3. 下线中...")
	if err = client.Logout(); err != nil {
		t.Logf("   下线失败: %v", err)
	} else {
		t.Log("   下线成功!")
	}

	t.Log("=== 工作流测试完成 ===")
}

// 基准测试：心跳性能
func BenchmarkHeartbeat(b *testing.B) {
	client, _ := sdk.NewClient(testConfig)
	_, err := client.Login(sdk.LoginRequest{
		Username: testUsername,
		Password: testPassword,
	})
	if err != nil {
		b.Skip("登录失败")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = client.Heartbeat()
	}
}

// 示例：完整使用流程
func Example_fullWorkflow() {
	// 创建客户端
	client, _ := sdk.NewClient(sdk.Config{
		BaseURL:   "https://app.l9.lc",
		AppKey:    "DoBHMi5mTLi8r4wg18MtmqPmZeK3eGuq",
		SecretKey: "UcS1CM6OnXUEbsbAugyZgxsBo3X1iJgV",
		Version:   "1.00",
	})

	// 登录
	userInfo, err := client.Login(sdk.LoginRequest{
		Username: testUsername,
		Password: testPassword,
	})
	if err != nil {
		fmt.Printf("登录失败: %v\n", err)
		return
	}
	fmt.Printf("登录成功: %s\n", userInfo.Username)

	// 启动心跳协程
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			if err := client.Heartbeat(); err != nil {
				fmt.Printf("心跳失败: %v\n", err)
				return
			}
		}
	}()

	// ... 业务逻辑 ...

	// 退出时下线
	_ = client.Logout()
}
