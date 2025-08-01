# GitHub Webhook Demo Project

A complete project for learning and demonstrating GitHub Webhook functionality, including Go server implementation and detailed educational documentation.

## 🚀 Project Overview

This project provides a complete GitHub Webhook receiving server that supports:

- Receiving and processing various GitHub events (Issues, Comments, Push, etc.)
- HMAC-SHA256 signature verification for security
- Support for local development and Render platform deployment
- Detailed logging and event processing
- Health check endpoint

## 📁 Project Structure

```
├── README.md              # Project documentation
├── webhook.md             # Complete Webhook tutorial
├── claude-code-intro.md   # Claude Code tool introduction
├── go.mod                 # Go module configuration (root)
├── render.yaml            # Render platform deployment config
└── webhook-demo/          # Webhook server implementation
    ├── go.mod            # Go module configuration
    └── server.go         # Main server code
```

## 🛠️ Quick Start

### Prerequisites

- Go 1.21 or higher
- Git

### Local Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd agent-test
   ```

2. **Run the server**
   ```bash
   cd webhook-demo
   go run server.go
   ```

3. **Test the server**
   ```bash
   curl http://localhost:8080/health
   ```

### Configure Webhook Secret (Recommended)

1. **Generate a secret**
   ```bash
   openssl rand -hex 20
   ```

2. **Set environment variable**
   ```bash
   export WEBHOOK_SECRET=your-generated-secret
   cd webhook-demo
   go run server.go
   ```

## 🌐 Deploy to Render

The project includes a `render.yaml` configuration file for one-click deployment to Render platform:

1. Connect your GitHub repository to Render
2. Select Web Service
3. Render will automatically use the configuration in `render.yaml`

After deployment, it will automatically:
- Set the `WEBHOOK_SECRET` environment variable
- Configure health checks
- Enable auto-deployment

## 📚 Usage Tutorial

For detailed usage tutorials, please see:
- [**webhook.md**](webhook.md) - Complete GitHub Webhook configuration and usage tutorial
- Includes detailed steps for tunneling, GitHub configuration, testing methods, etc.

## 🎯 Supported Events

The server can handle the following GitHub events:

- **ping** - Webhook connection test
- **issues** - Issue creation, editing, closing, etc.
- **issue_comment** - Issue comments (supports command detection)
- **push** - Code pushes
- Other events are also logged

## 🔧 API Endpoints

- `POST /webhook` - Webhook event receiving endpoint
- `GET /health` - Health check endpoint

## 🔒 Security Features

- HMAC-SHA256 signature verification
- GitHub Webhook Secret support
- Detailed security logging
- Input validation and error handling

## 📖 Related Documentation

- [Claude Code Introduction](claude-code-intro.md) - AI programming assistant tool introduction
- [Webhook Tutorial](webhook.md) - Complete tutorial from scratch

## 🤝 Contributing

Issues and Pull Requests are welcome to improve this project!

## 📄 License

This project is for learning and demonstration purposes only.