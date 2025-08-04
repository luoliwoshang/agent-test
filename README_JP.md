# GitHub Webhook デモプロジェクト

GitHub Webhook機能の学習とデモンストレーションのための完全なプロジェクトです。Goサーバー実装と詳細な教育文書が含まれています。

## 🚀 プロジェクト概要

このプロジェクトは完全なGitHub Webhook受信サーバーを提供し、以下をサポートします：

- 各種GitHub イベント（Issues、Comments、Push等）の受信と処理
- セキュリティを確保するHMAC-SHA256署名検証
- ローカル開発とRenderプラットフォームデプロイのサポート
- 詳細なログ出力とイベント処理
- ヘルスチェックエンドポイント

## 📁 プロジェクト構造

```
├── README.md              # プロジェクト説明文書
├── webhook.md             # 完全なWebhook学習チュートリアル
├── claude-code-intro.md   # Claude Code ツール紹介
├── go.mod                 # Goモジュール設定（ルートディレクトリ）
├── render.yaml            # Renderプラットフォームデプロイ設定
└── webhook-demo/          # Webhookサーバー実装
    ├── go.mod            # Goモジュール設定
    └── server.go         # メインサーバーコード
```

## 🛠️ クイックスタート

### 前提条件

- Go 1.21以上
- Git

### ローカル実行

1. **プロジェクトのクローン**
   ```bash
   git clone <repository-url>
   cd agent-test
   ```

2. **サーバーの実行**
   ```bash
   cd webhook-demo
   go run server.go
   ```

3. **サーバーのテスト**
   ```bash
   curl http://localhost:8080/health
   ```

### Webhook Secretの設定（推奨）

1. **秘密鍵の生成**
   ```bash
   openssl rand -hex 20
   ```

2. **環境変数の設定**
   ```bash
   export WEBHOOK_SECRET=your-generated-secret
   cd webhook-demo
   go run server.go
   ```

## 🌐 Renderへのデプロイ

プロジェクトには`render.yaml`設定ファイルが含まれており、Renderプラットフォームへのワンクリックデプロイが可能です：

1. GitHubリポジトリをRenderに接続
2. Web Serviceを選択
3. Renderが自動的に`render.yaml`の設定を使用

デプロイ後、以下が自動的に行われます：
- `WEBHOOK_SECRET`環境変数の設定
- ヘルスチェックの設定
- 自動デプロイの有効化

## 📚 使用チュートリアル

詳細な使用チュートリアルは以下をご覧ください：
- [**webhook.md**](webhook.md) - 完全なGitHub Webhook設定と使用チュートリアル
- 内部ネットワーク貫通、GitHub設定、テスト方法等の詳細手順を含む

## 🎯 サポートされるイベント

サーバーは以下のGitHubイベントを処理できます：

- **ping** - Webhook接続テスト
- **issues** - Issue作成、編集、クローズ等
- **issue_comment** - Issueコメント（コマンド検出サポート）
- **push** - コードプッシュ
- その他のイベントも記録されます

## 🔧 APIエンドポイント

- `POST /webhook` - Webhookイベント受信エンドポイント
- `GET /health` - ヘルスチェックエンドポイント

## 🔒 セキュリティ機能

- HMAC-SHA256署名検証
- GitHub Webhook Secretサポート
- 詳細なセキュリティログ記録
- 入力検証とエラーハンドリング

## 📖 関連文書

- [Claude Code 紹介](claude-code-intro.md) - AIプログラミングアシスタントツール紹介
- [Webhook チュートリアル](webhook.md) - ゼロから始める完全教学

## 🤝 貢献

このプロジェクトの改善のためのIssueやPull Requestの提出を歓迎します！

## 📄 ライセンス

このプロジェクトは学習とデモンストレーション目的のみで使用されます。