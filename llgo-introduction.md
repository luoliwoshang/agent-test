**LLGo is an innovative LLVM-based Go compiler that enables seamless integration between Go and C/C++, Python, JavaScript ecosystems.**

# LLGo - 基于 LLVM 的 Go 编译器

## 🚀 项目简介

LLGo 是一个创新的基于 LLVM 的 Go 编译器，旨在打破语言边界，实现 Go 语言与 C/C++、Python、JavaScript 等生态系统的无缝集成。该项目由 Go+ 团队开发，为 Go 开发者提供了更强大的跨语言编程能力。

## ✨ 核心特性

### 🔗 跨语言集成
- **直接调用 C/C++ 标准库**：无需复杂的 CGO 绑定，直接使用 C 函数
- **Python 库支持**：可以在 Go 代码中直接导入和使用 Python 库
- **JavaScript 互操作**：支持与 JavaScript 生态的集成
- **完整 Go 语法兼容**：保持 Go 语言的原生语法和特性

### ⚡ 技术优势
- **LLVM 后端**：使用现代化的 LLVM 编译器基础设施
- **自动垃圾回收**：默认集成 bdwgc 垃圾收集器
- **WebAssembly 支持**：可编译生成 WebAssembly 模块
- **高性能编译**：利用 LLVM 的优化能力提供高效代码生成

## 🏗️ 技术架构

LLGo 采用模块化设计，主要包含以下核心组件：

- **ssa 模块**：负责生成 LLVM 中间表示（IR）文件
- **cl 模块**：将 Go 包转换为 LLVM IR
- **internal/build 模块**：管理整个编译流程和构建过程

## 🛠️ 环境要求

### 基础依赖
- **Go 1.21+**：Go 语言运行时环境
- **LLVM 18**：LLVM 编译器基础设施
- **Clang 18**：C/C++ 编译器
- **Python 3.12+**（可选）：用于 Python 集成功能

### 平台支持
- **macOS**：通过 Homebrew 安装依赖
- **Linux**：支持主流发行版的包管理器
- **源码编译**：支持从源代码自定义编译安装

## 💡 使用示例

### 调用 C 标准库

```go
package main

import "github.com/goplus/lib/c"

func main() {
    c.Printf(c.Str("Hello from C library!\n"))
    
    // 使用 C 数学函数
    result := c.Sqrt(16.0)
    c.Printf(c.Str("sqrt(16) = %.2f\n"), result)
}
```

### 集成 Python 库

```go
package main

import (
    "github.com/goplus/lib/py"
    "github.com/goplus/lib/py/math"
)

func main() {
    // 使用 Python 数学库
    result := math.Sin(py.Float(3.14159 / 2))
    py.Print("sin(π/2) =", result)
    
    // 调用 Python 内置函数
    py.Print("Hello from Python!")
}
```

### WebAssembly 编译

```bash
# 编译为 WebAssembly
llgo build -target wasm ./your-go-program.go
```

## 🎯 应用场景

### 游戏开发
- 利用 C/C++ 游戏引擎的性能优势
- 结合 Go 的并发特性和简洁语法
- 支持跨平台部署

### 人工智能与数据科学
- 直接使用 Python 的 AI/ML 库（如 NumPy、TensorFlow）
- 享受 Go 的类型安全和性能优势
- 简化部署和分发流程

### Web 开发
- 编译为 WebAssembly 在浏览器中运行
- 复用服务端 Go 代码逻辑
- 提供近似原生的性能体验

### 系统编程
- 直接调用系统 C 库
- 保持 Go 的内存安全特性
- 简化与现有 C/C++ 代码库的集成

## 🚀 快速开始

### 1. 安装依赖

**macOS (使用 Homebrew)：**
```bash
brew install llvm@18 python@3.12
```

**Ubuntu/Debian：**
```bash
sudo apt-get install llvm-18 clang-18 python3.12
```

### 2. 安装 LLGo

```bash
go install github.com/goplus/llgo/cmd/llgo@latest
```

### 3. 编译运行

```bash
# 创建示例程序
echo 'package main
import "github.com/goplus/lib/c"
func main() { c.Printf(c.Str("Hello LLGo!\n")) }' > hello.go

# 编译并运行
llgo run hello.go
```

## 🔧 进阶配置

### 自定义编译选项

```bash
# 指定优化级别
llgo build -O2 ./program.go

# 生成调试信息
llgo build -g ./program.go

# 静态链接
llgo build -static ./program.go
```

### 集成现有项目

```go
// go.mod
module your-project

require github.com/goplus/llgo v0.9.0
```

## 🌟 项目优势

1. **无缝集成**：消除语言边界，直接使用各语言生态库
2. **性能优越**：基于 LLVM 的优化编译，提供卓越性能
3. **开发友好**：保持 Go 语言简洁优雅的开发体验
4. **生态丰富**：可访问 C/C++、Python、JavaScript 的庞大生态
5. **部署灵活**：支持传统二进制、WebAssembly 等多种部署方式

## 📚 相关资源

- **GitHub 仓库**：https://github.com/goplus/llgo
- **官方文档**：查看项目 Wiki 获取详细文档
- **社区支持**：通过 GitHub Issues 获取技术支持
- **Go+ 官网**：了解更多 Go+ 生态项目

## 🤝 贡献指南

LLGo 是一个开源项目，欢迎社区贡献：

1. **报告问题**：通过 GitHub Issues 报告 bug 或提出建议
2. **提交代码**：遵循项目的代码规范提交 Pull Request  
3. **完善文档**：帮助改进项目文档和示例
4. **测试反馈**：在不同平台和场景下测试并提供反馈

---

LLGo 代表了 Go 语言发展的新方向，通过打破语言壁垒，为开发者提供了前所未有的编程体验。无论您是系统开发者、游戏开发者还是数据科学家，LLGo 都能为您的项目带来新的可能性。