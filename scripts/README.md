# 脚本工具

## validate-links.py

文档链接安全验证脚本，用于CI流程中验证提交的文档是否包含安全的链接。

### 功能特性

- 🔍 **链接提取**：从Markdown文档中提取所有类型的链接
- 🛡️ **安全检查**：检查链接是否安全，防止恶意链接
- 📋 **详细报告**：提供详细的安全问题报告
- 🚀 **CI集成**：专为CI/CD流程设计

### 安全检查项目

1. **危险协议检测**
   - `javascript:` - 可能执行恶意脚本
   - `ftp:` - 不安全的文件传输
   - `file:` - 本地文件访问
   - `data:` - 可能包含恶意数据

2. **黑名单域名**
   - 短链接服务：`bit.ly`, `tinyurl.com`, `goo.gl`
   - 已知恶意域名

3. **可疑模式**
   - IP地址直接访问
   - 过长域名
   - 随机字符域名

### 使用方法

```bash
# 验证当前目录下的所有markdown文件
python3 scripts/validate-links.py

# 在CI中使用
chmod +x scripts/validate-links.py
./scripts/validate-links.py
```

### 退出码

- `0` - 所有链接都安全
- `1` - 发现安全问题

### 配置

可以通过修改脚本中的以下变量来调整检查规则：

- `DANGEROUS_DOMAINS` - 危险域名黑名单
- `DANGEROUS_PROTOCOLS` - 危险协议列表
- `SAFE_DOMAINS` - 安全域名白名单