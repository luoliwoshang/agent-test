# 🎯 GitHub Webhook 完整教学教程

## 📚 第一步：理解 Webhook

**Webhook 就像快递员：**
- 普通方式：你每隔一段时间去门口看有没有快递（轮询）
- Webhook 方式：快递到了，快递员主动按门铃通知你（推送）

**GitHub Webhook 工作原理：**
```
GitHub 上发生事件 → GitHub 立即发送 HTTP 请求 → 你的服务器接收处理
```

## 🔧 第二步：准备本地服务器

假设你已经有了一个可以接收 webhook 的服务器运行在 `localhost:8080`。

**测试服务器是否正常：**
```bash
curl http://localhost:8080/health
```
应该返回：`Webhook server is running!`

## 🌐 第三步：内网穿透设置

**为什么需要内网穿透？**
- GitHub 在互联网上，你的服务器在本地
- GitHub 无法直接访问你的 `localhost:8080`
- 需要一个工具把你的本地服务暴露到互联网

### 方法1：使用 ngrok（推荐）

**1. 安装 ngrok**
```bash
# macOS
brew install ngrok

# Windows (下载安装包)
# 去 https://ngrok.com/download 下载

# Linux
wget https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip
unzip ngrok-stable-linux-amd64.zip
```

**2. 创建隧道**
```bash
ngrok http 8080
```

**3. 你会看到这样的输出：**
```
ngrok by @inconshreveable

Session Status                online
Account                       你的账号
Version                       2.x.x
Region                        United States (us)
Web Interface                 http://127.0.0.1:4040
Forwarding                    http://abc123.ngrok.io -> http://localhost:8080
Forwarding                    https://abc123.ngrok.io -> http://localhost:8080
```

**重要：记住这个 URL：`https://abc123.ngrok.io`**

**4. 测试外网访问**
```bash
curl https://abc123.ngrok.io/health
```

### 方法2：其他工具

**localtunnel:**
```bash
npm install -g localtunnel
lt --port 8080
```

**serveo:**
```bash
ssh -R 80:localhost:8080 serveo.net
```

## ⚙️ 第四步：GitHub Webhook 配置

**1. 进入你的 GitHub 仓库设置**
- 去：`https://github.com/你的用户名/你的仓库名/settings/hooks`
- 点击 **"Add webhook"**

**2. 填写配置信息**

| 字段 | 值 | 说明 |
|------|-----|------|
| **Payload URL** | `https://abc123.ngrok.io/webhook` | 你的 ngrok 地址 + `/webhook` |
| **Content type** | `application/json` | JSON 格式数据 |
| **Secret** | 留空（暂时） | 安全验证，后面再讲 |
| **SSL verification** | ✅ Enable | 使用 HTTPS |

**3. 选择触发事件**

有三种选择：
- **Just the push event** - 只有代码推送
- **Send me everything** - 所有事件（学习阶段推荐）
- **Let me select individual events** - 自定义选择

**推荐选择具体事件：**
- ✅ Issues
- ✅ Issue comments  
- ✅ Pull requests
- ✅ Pull request reviews
- ✅ Pushes

**4. 点击 "Add webhook"**

## 🧪 第五步：测试 Webhook

### 自动测试：Ping 事件

配置完成后，GitHub 立即发送一个 `ping` 事件测试连接。

**你的服务器日志应该显示：**
```
收到 webhook 请求：POST /webhook
事件类型: ping
🏓 收到 ping 事件 - webhook 配置成功！
```

**在 GitHub 页面上也能看到：**
- 绿色勾号 ✅ = 成功
- 红色叉号 ❌ = 失败

### 手动测试：触发各种事件

**1. Issue 事件测试**
- 在仓库中创建一个新 Issue
- 标题：`测试 webhook`
- 内容：`这是一个测试 Issue`

**服务器日志会显示：**
```
事件类型: issues
📝 Issue 事件: opened
  Issue #1: 测试 webhook
```

**2. Issue 评论测试**
- 在刚创建的 Issue 中添加评论：`/hello`

**服务器日志会显示：**
```
事件类型: issue_comment
💬 Issue 评论事件: created
  Issue #1 收到评论: /hello
🎉 检测到 /hello 命令!
```

**3. Push 事件测试**
- 在仓库中创建一个文件：`echo "test" > test.txt`
- 提交并推送：`git add . && git commit -m "test" && git push`

**服务器日志会显示：**
```
事件类型: push
🚀 Push 事件 to 你的用户名/仓库名
```

## 📋 第六步：理解 Payload 数据

每个 webhook 请求都包含详细的 JSON 数据。

### Issue Comment Payload 示例

```json
{
  "action": "created",
  "issue": {
    "number": 1,
    "title": "测试 webhook",
    "body": "这是一个测试 Issue",
    "user": {
      "login": "你的用户名"
    }
  },
  "comment": {
    "body": "/hello",
    "user": {
      "login": "你的用户名"
    }
  },
  "repository": {
    "name": "agent-test",
    "full_name": "你的用户名/agent-test"
  },
  "sender": {
    "login": "你的用户名"
  }
}
```

### 关键字段说明

- `action`: 事件动作（created、opened、closed 等）
- `issue`: Issue 详细信息
- `comment`: 评论内容
- `repository`: 仓库信息  
- `sender`: 触发事件的用户

## 🔍 第七步：调试技巧

**1. 查看 ngrok 请求详情**
- 打开：`http://127.0.0.1:4040`
- 可以看到所有请求的详细信息

**2. GitHub Webhook 页面**
- 在 webhook 配置页面点击 webhook
- 可以看到所有请求的状态和响应

**3. 重新发送测试**
- 在 GitHub webhook 详情页面
- 点击任意一个请求旁边的 "Redeliver"
- 可以重新发送该请求用于调试

## 🔐 第八步：添加 Webhook Secret 安全验证

**为什么需要 Secret 验证？**
- 确保请求真的来自 GitHub，而不是恶意攻击者
- 防止他人伪造 webhook 请求
- 这是生产环境的必备安全措施

### 1. 生成 Secret

**生成一个强密码作为 Secret：**
```bash
# 使用 openssl 生成随机字符串
openssl rand -hex 20

# 或者使用 Python
python3 -c "import secrets; print(secrets.token_hex(20))"

# 示例输出: a1b2c3d4e5f6789012345678901234567890abcd
```

**记住这个 Secret，我们需要在两个地方使用它。**

### 2. 设置环境变量

**在启动服务器前设置环境变量：**
```bash
# 设置你的 secret（替换成你生成的值）
export WEBHOOK_SECRET=a1b2c3d4e5f6789012345678901234567890abcd

# 启动服务器
cd webhook-demo
go run server.go
```

**现在服务器会显示：**
```
🚀 启动 Webhook 演示服务器...
📡 监听端口: 8080
🔗 Webhook URL: http://localhost:8080/webhook
🏥 健康检查: http://localhost:8080/health
🔐 已启用 Webhook 签名验证
```

### 3. 更新 GitHub Webhook 配置

**回到 GitHub webhook 配置页面：**
1. 进入：`https://github.com/你的用户名/你的仓库名/settings/hooks`
2. 点击你之前创建的 webhook
3. 点击 **"Edit"**
4. 在 **"Secret"** 字段填入你生成的 secret
5. 点击 **"Update webhook"**

### 4. 验证 Secret 功能

**测试没有 Secret 的情况：**
```bash
# 先停止服务器，然后不设置环境变量直接启动
go run server.go
```

服务器会显示警告：
```
⚠️  警告: 未设置 WEBHOOK_SECRET 环境变量
   建议设置: export WEBHOOK_SECRET=your-secret-key
   这样可以启用签名验证以确保安全性
```

**测试错误的 Secret：**
```bash
# 设置错误的 secret
export WEBHOOK_SECRET=wrong-secret
go run server.go
```

触发一个 webhook 事件，服务器日志会显示：
```
❌ Webhook 签名验证失败，拒绝请求
❌ 签名验证失败
期望: abc123...
实际: def456...
```

**测试正确的 Secret：**
```bash
# 设置正确的 secret
export WEBHOOK_SECRET=a1b2c3d4e5f6789012345678901234567890abcd
go run server.go
```

触发事件后，服务器日志会显示：
```
✅ 签名验证成功
🏓 收到 ping 事件 - webhook 配置成功！
```

### 5. Secret 验证原理说明

**GitHub 如何生成签名：**
1. 使用你设置的 Secret 作为密钥
2. 对请求体内容计算 HMAC-SHA256 哈希
3. 在请求头中发送：`X-Hub-Signature-256: sha256=计算的哈希值`

**服务器如何验证：**
1. 获取请求头中的签名
2. 使用相同的 Secret 对请求体计算 HMAC-SHA256
3. 比较两个哈希值是否相同
4. 相同则验证通过，不同则拒绝请求

**安全注意事项：**
- Secret 要足够长和随机（至少 20 字符）
- 不要在代码中硬编码 Secret
- 定期更换 Secret
- 使用环境变量或安全的配置管理系统

## 🎉 完成！

现在你已经成功搭建了一个完整且安全的 webhook 接收系统：

✅ 本地服务器运行  
✅ 内网穿透配置  
✅ GitHub webhook 配置  
✅ 事件接收和处理  
✅ 调试和测试  
✅ **安全验证（Webhook Secret）**

**下一步可以做什么？**
- 添加更复杂的事件处理逻辑
- 集成 AI 服务进行代码生成
- 自动创建 PR 和提交代码
- 添加数据库存储 webhook 事件

这就是一个完整且安全的 webhook 教学流程！