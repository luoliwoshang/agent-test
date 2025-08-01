#!/usr/bin/env python3
"""
文档链接安全验证脚本
检查文档中的链接是否安全，防止恶意链接入库
"""

import re
import sys
import os
import glob
from urllib.parse import urlparse
from pathlib import Path

# 危险域名黑名单
DANGEROUS_DOMAINS = {
    # 恶意软件/钓鱼网站常用域名
    'bit.ly', 'tinyurl.com', 'goo.gl', 't.co',  # 短链接服务
    'malware.com', 'phishing.com', 'virus.com',  # 明显恶意域名
    # 可以根据需要添加更多
}

# 可疑协议
DANGEROUS_PROTOCOLS = {
    'ftp', 'file', 'javascript', 'data', 'vbscript'
}

# 允许的安全域名
SAFE_DOMAINS = {
    'github.com', 'githubusercontent.com',
    'docs.github.com', 'help.github.com',
    'stackoverflow.com', 'developer.mozilla.org',
    'w3.org', 'ietf.org', 'rfc-editor.org',
    'nodejs.org', 'python.org', 'golang.org',
    'microsoft.com', 'google.com', 'anthropic.com',
    'ngrok.com', 'render.com', 'localhost',
    'serveo.net', 'localtunnel.me',
    'example.com', 'example.org',  # 示例域名
}

# 本地链接模式
LOCAL_PATTERNS = [
    r'^#',           # 锚点链接
    r'^\./',         # 相对路径
    r'^\.\.',        # 上级目录
    r'^/',           # 绝对路径
    r'^[^/]+\.md$',  # 同目录markdown文件
]

def extract_links_from_markdown(content):
    """从markdown内容中提取所有链接"""
    links = []
    
    # Markdown链接格式: [text](url)
    markdown_links = re.findall(r'\[([^\]]*)\]\(([^)]+)\)', content)
    for text, url in markdown_links:
        links.append({'text': text, 'url': url, 'type': 'markdown'})
    
    # HTML链接格式: <a href="url">text</a>
    html_links = re.findall(r'<a[^>]+href=["\']([^"\']+)["\'][^>]*>([^<]*)</a>', content, re.IGNORECASE)
    for url, text in html_links:
        links.append({'text': text, 'url': url, 'type': 'html'})
    
    # 纯URL (http/https开头的独立URL)
    url_pattern = r'https?://[^\s<>"\'\[\]()]+[^\s<>"\'\[\]().,;:]'
    pure_urls = re.findall(url_pattern, content)
    for url in pure_urls:
        # 避免重复（已经在markdown或html链接中的URL）
        if not any(link['url'] == url for link in links):
            links.append({'text': url, 'url': url, 'type': 'plain'})
    
    return links

def is_local_link(url):
    """检查是否为本地链接"""
    for pattern in LOCAL_PATTERNS:
        if re.match(pattern, url):
            return True
    return False

def validate_url(url):
    """验证单个URL的安全性"""
    issues = []
    
    # 跳过本地链接
    if is_local_link(url):
        return issues
    
    try:
        parsed = urlparse(url)
        
        # 检查协议
        if parsed.scheme.lower() in DANGEROUS_PROTOCOLS:
            issues.append(f"危险协议: {parsed.scheme}")
        
        # 检查域名
        domain = parsed.netloc.lower()
        if domain:
            # 移除端口号
            domain = domain.split(':')[0]
            
            # 检查是否在黑名单中
            if domain in DANGEROUS_DOMAINS:
                issues.append(f"黑名单域名: {domain}")
            
            # 检查可疑模式
            if re.search(r'\d+\.\d+\.\d+\.\d+', domain):  # IP地址
                if not domain.startswith('127.') and not domain.startswith('192.168.'):
                    issues.append(f"使用IP地址而非域名: {domain}")
            
            # 检查过长域名
            if len(domain) > 100:
                issues.append("域名过长，可能是恶意域名")
            
            # 检查随机字符域名
            if re.match(r'^[a-z0-9]{10,}\.', domain):
                issues.append("疑似随机生成的域名")
    
    except Exception as e:
        issues.append(f"URL解析错误: {e}")
    
    return issues

def validate_file(file_path):
    """验证单个文件中的链接"""
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            content = f.read()
    except Exception as e:
        return [{'file': file_path, 'error': f"文件读取错误: {e}"}]
    
    links = extract_links_from_markdown(content)
    issues = []
    
    for link in links:
        url = link['url']
        link_issues = validate_url(url)
        
        if link_issues:
            issues.append({
                'file': file_path,
                'text': link['text'],
                'url': url,
                'type': link['type'],
                'issues': link_issues
            })
    
    return issues

def main():
    """主函数"""
    print("🔍 开始验证文档链接安全性...")
    
    # 查找所有markdown文件
    md_files = []
    for pattern in ['**/*.md', '*.md']:
        md_files.extend(glob.glob(pattern, recursive=True))
    
    if not md_files:
        print("❌ 未找到任何markdown文件")
        return 1
    
    print(f"📝 找到 {len(md_files)} 个markdown文件")
    
    all_issues = []
    
    for file_path in md_files:
        print(f"  检查: {file_path}")
        issues = validate_file(file_path)
        all_issues.extend(issues)
    
    # 输出结果
    if not all_issues:
        print("✅ 所有链接都是安全的！")
        return 0
    
    print(f"\n⚠️  发现 {len(all_issues)} 个潜在问题:")
    
    current_file = None
    for issue in all_issues:
        if 'error' in issue:
            print(f"\n❌ {issue['file']}: {issue['error']}")
            continue
            
        if issue['file'] != current_file:
            current_file = issue['file']
            print(f"\n📄 {current_file}:")
        
        print(f"  🔗 [{issue['text']}]({issue['url']})")
        print(f"     类型: {issue['type']}")
        for problem in issue['issues']:
            print(f"     ⚠️  {problem}")
    
    print(f"\n❌ 验证失败: 发现 {len(all_issues)} 个安全问题")
    return 1

if __name__ == '__main__':
    sys.exit(main())