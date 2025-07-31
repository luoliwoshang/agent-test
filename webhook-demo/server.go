package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// WebhookPayload 简化的 webhook 数据结构
type WebhookPayload struct {
	Action     string `json:"action,omitempty"`
	Repository struct {
		Name     string `json:"name"`
		FullName string `json:"full_name"`
	} `json:"repository"`
	Sender struct {
		Login string `json:"login"`
	} `json:"sender"`
	Issue *struct {
		Number int    `json:"number"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	} `json:"issue,omitempty"`
	Comment *struct {
		Body string `json:"body"`
	} `json:"comment,omitempty"`
}

// verifyWebhookSignature 验证 webhook 签名
func verifyWebhookSignature(payload []byte, signature string, secret string) bool {
	if secret == "" {
		log.Println("⚠️  警告: 未设置 WEBHOOK_SECRET，跳过签名验证")
		return true
	}

	if signature == "" {
		log.Println("❌ 缺少签名头")
		return false
	}

	// GitHub 发送的签名格式是 "sha256=..."
	if !strings.HasPrefix(signature, "sha256=") {
		log.Println("❌ 签名格式错误")
		return false
	}

	// 提取实际的签名
	expectedSignature := signature[7:] // 去掉 "sha256=" 前缀

	// 使用 secret 计算 HMAC-SHA256
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	actualSignature := hex.EncodeToString(mac.Sum(nil))

	// 比较签名
	if hmac.Equal([]byte(expectedSignature), []byte(actualSignature)) {
		log.Println("✅ 签名验证成功")
		return true
	}

	log.Println("❌ 签名验证失败")
	log.Printf("期望: %s", expectedSignature)
	log.Printf("实际: %s", actualSignature)
	return false
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	// 记录请求信息
	log.Printf("收到 webhook 请求：%s %s", r.Method, r.URL.Path)
	log.Printf("Headers: %v", r.Header)

	// 读取请求体
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("读取请求体失败: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// 验证 webhook 签名
	signature := r.Header.Get("X-Hub-Signature-256")
	secret := os.Getenv("WEBHOOK_SECRET")
	
	if !verifyWebhookSignature(body, signature, secret) {
		log.Println("❌ Webhook 签名验证失败，拒绝请求")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 获取事件类型
	eventType := r.Header.Get("X-GitHub-Event")
	if eventType == "" {
		log.Println("缺少 X-GitHub-Event header")
		http.Error(w, "Missing event type", http.StatusBadRequest)
		return
	}

	log.Printf("事件类型: %s", eventType)
	log.Printf("请求体大小: %d bytes", len(body))

	// 解析 JSON
	var payload WebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Printf("JSON 解析失败: %v", err)
		// 即使解析失败也要返回成功，避免 GitHub 重试
		w.WriteHeader(http.StatusOK)
		return
	}

	// 根据事件类型处理
	switch eventType {
	case "ping":
		log.Println("🏓 收到 ping 事件 - webhook 配置成功！")

	case "issues":
		log.Printf("📝 Issue 事件: %s", payload.Action)
		if payload.Issue != nil {
			log.Printf("  Issue #%d: %s", payload.Issue.Number, payload.Issue.Title)
		}

	case "issue_comment":
		log.Printf("💬 Issue 评论事件: %s", payload.Action)
		if payload.Issue != nil && payload.Comment != nil {
			log.Printf("  Issue #%d 收到评论: %s", payload.Issue.Number, payload.Comment.Body)

			// 检查是否包含特殊命令
			if payload.Comment.Body == "/hello" {
				log.Println("🎉 检测到 /hello 命令!")
				// 这里可以执行特定逻辑
			}
		}

	case "push":
		log.Printf("🚀 Push 事件 to %s", payload.Repository.FullName)

	default:
		log.Printf("ℹ️  未处理的事件类型: %s", eventType)
	}

	log.Printf("事件来源: %s (用户: %s)", payload.Repository.FullName, payload.Sender.Login)
	log.Println("---")

	// 返回成功响应
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received successfully"))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Webhook server is running! Time: %s", time.Now().Format("2006-01-02 15:04:05"))))
}

func main() {
	// 获取端口配置，支持 Render 和本地开发
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // 本地开发默认端口
	}

	log.Println("🚀 启动 Webhook 演示服务器...")
	log.Printf("📡 监听端口: %s", port)
	
	// 根据环境显示不同的 URL
	if os.Getenv("RENDER") != "" {
		log.Println("🌐 运行环境: Render Cloud")
		log.Println("🔗 Webhook URL: https://your-app.onrender.com/webhook")
		log.Println("🏥 健康检查: https://your-app.onrender.com/health")
	} else {
		log.Printf("🔗 Webhook URL: http://localhost:%s/webhook", port)
		log.Printf("🏥 健康检查: http://localhost:%s/health", port)
	}
	
	// 检查是否设置了 WEBHOOK_SECRET
	secret := os.Getenv("WEBHOOK_SECRET")
	if secret == "" {
		log.Println("⚠️  警告: 未设置 WEBHOOK_SECRET 环境变量")
		log.Println("   建议设置: export WEBHOOK_SECRET=your-secret-key")
		log.Println("   这样可以启用签名验证以确保安全性")
	} else {
		log.Println("🔐 已启用 Webhook 签名验证")
	}
	
	log.Println()

	http.HandleFunc("/webhook", webhookHandler)
	http.HandleFunc("/health", healthHandler)

	log.Printf("🎯 服务器启动成功，监听端口 %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
