package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	log.Println("🚀 启动 Webhook 演示服务器...")
	log.Println("📡 监听端口: 8080")
	log.Println("🔗 Webhook URL: http://localhost:8080/webhook")
	log.Println("🏥 健康检查: http://localhost:8080/health")
	log.Println()

	http.HandleFunc("/webhook", webhookHandler)
	http.HandleFunc("/health", healthHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
