# GitHub Webhook 演示项目

一个用于学习和演示 GitHub Webhook 功能的完整项目，包含 Go 服务器实现和详细的教学文档。

## 🚀 项目简介

本项目提供了一个完整的 GitHub Webhook 接收服务器，支持：

- 接收和处理各种 GitHub 事件（Issues、Comments、Push 等）
- HMAC-SHA256 签名验证确保安全
- 支持本地开发和 Render 平台部署
- 详细的日志输出和事件处理
- 健康检查端点

## 📁 项目结构

```
├── README.md              # 项目说明文档
├── webhook.md             # 完整的 Webhook 教学教程
├── claude-code-intro.md   # Claude Code 工具简介
├── go.mod                 # Go 模块配置（根目录）
├── render.yaml            # Render 平台部署配置
└── webhook-demo/          # Webhook 服务器实现
    ├── go.mod            # Go 模块配置
    └── server.go         # 主服务器代码
```

## 🛠️ 快速开始

### 前置要求

- Go 1.21 或更高版本
- Git

### 本地运行

1. **克隆项目**
   ```bash
   git clone <repository-url>
   cd agent-test
   ```

2. **运行服务器**
   ```bash
   cd webhook-demo
   go run server.go
   ```

3. **测试服务器**
   ```bash
   curl http://localhost:8080/health
   ```

### 使用主机配置测试

项目支持读取主机配置进行测试：

1. **使用主机环境变量**
   ```bash
   # 服务器会自动读取主机的环境变量配置
   # 包括 PORT、WEBHOOK_SECRET、RENDER 等
   cd webhook-demo
   go run server.go
   ```

2. **主机配置检测**
   - 自动检测 `PORT` 环境变量（默认 8080）
   - 自动检测 `WEBHOOK_SECRET` 进行签名验证
   - 自动检测 `RENDER` 环境标识切换运行模式
   - 显示对应环境的访问 URL 和配置信息

3. **配置验证测试**
   ```bash
   # 测试不同配置场景
   PORT=3000 go run server.go                    # 自定义端口
   WEBHOOK_SECRET=test123 go run server.go       # 启用签名验证
   RENDER=true go run server.go                  # 模拟云环境
   ```

### 配置 Webhook Secret（推荐）

1. **生成密钥**
   ```bash
   openssl rand -hex 20
   ```

2. **设置环境变量**
   ```bash
   export WEBHOOK_SECRET=your-generated-secret
   cd webhook-demo
   go run server.go
   ```

## 🌐 部署到 Render

项目包含了 `render.yaml` 配置文件，可以一键部署到 Render 平台：

1. 连接 GitHub 仓库到 Render
2. 选择 Web Service
3. Render 会自动使用 `render.yaml` 中的配置

部署后会自动：
- 设置 `WEBHOOK_SECRET` 环境变量
- 配置健康检查
- 启用自动部署

## 📚 使用教程

详细的使用教程请查看：
- [**webhook.md**](webhook.md) - 完整的 GitHub Webhook 配置和使用教程
- 包含内网穿透、GitHub 配置、测试方法等详细步骤

## 🎯 支持的事件

服务器可以处理以下 GitHub 事件：

- **ping** - Webhook 连接测试
- **issues** - Issue 创建、编辑、关闭等
- **issue_comment** - Issue 评论（支持命令检测）
- **push** - 代码推送
- 其他事件也会被记录

## 🔧 API 端点

- `POST /webhook` - Webhook 事件接收端点
- `GET /health` - 健康检查端点

## 🔒 安全特性

- HMAC-SHA256 签名验证
- 支持 GitHub Webhook Secret
- 详细的安全日志记录
- 输入验证和错误处理

## 📖 相关文档

- [Claude Code 简介](claude-code-intro.md) - AI 编程助手工具介绍
- [Webhook 教程](webhook.md) - 从零开始的完整教学

## 🤝 贡献

欢迎提交 Issue 和 Pull Request 来改进这个项目！

## 📄 许可证

本项目仅用于学习和演示目的。
