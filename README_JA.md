# GitHub Webhook デモプロジェクト

GitHub Webhook機能の学習とデモンストレーション用の完全なプロジェクトです。Goサーバー実装と詳細なチュートリアルドキュメントが含まれています。

## 🚀 プロジェクト概要

このプロジェクトは以下をサポートする完全なGitHub Webhook受信サーバーを提供します：

- 各種GitHubイベント（Issues、Comments、Pushなど）の受信と処理
- セキュリティのためのHMAC-SHA256署名検証
- ローカル開発とRenderプラットフォームデプロイメントのサポート
- 詳細なログ出力とイベント処理
- ヘルスチェックエンドポイント

## 📁 プロジェクト構造

```
├── README.md              # プロジェクトドキュメント
├── webhook.md             # 完全なWebhookチュートリアル
├── claude-code-intro.md   # Claude Codeツール紹介
├── go.mod                 # Goモジュール設定（ルート）
├── render.yaml            # Renderプラットフォームデプロイ設定
└── webhook-demo/          # Webhookサーバー実装
    ├── go.mod            # Goモジュール設定
    └── server.go         # メインサーバーコード
```

## 🛠️ クイックスタート

### 前提条件

- Go 1.21以上
- Git

### ローカル開発

1. **プロジェクトをクローン**
   ```bash
   git clone <repository-url>
   cd agent-test
   ```

2. **サーバーを実行**
   ```bash
   cd webhook-demo
   go run server.go
   ```

3. **サーバーをテスト**
   ```bash
   curl http://localhost:8080/health
   ```

### Webhook Secretの設定（推奨）

1. **シークレットを生成**
   ```bash
   openssl rand -hex 20
   ```

2. **環境変数を設定**
   ```bash
   export WEBHOOK_SECRET=your-generated-secret
   cd webhook-demo
   go run server.go
   ```

## 🌐 Renderへのデプロイ

プロジェクトには`render.yaml`設定ファイルが含まれており、Renderプラットフォームにワンクリックでデプロイできます：

1. GitHubリポジトリをRenderに接続
2. Web Serviceを選択
3. Renderが自動的に`render.yaml`の設定を使用

デプロイ後、自動的に以下が設定されます：
- `WEBHOOK_SECRET`環境変数の設定
- ヘルスチェックの設定
- 自動デプロイの有効化

## 📚 使用方法チュートリアル

詳細な使用方法チュートリアルについては、以下をご確認ください：
- [**webhook.md**](webhook.md) - 完全なGitHub Webhook設定と使用方法チュートリアル
- トンネリング、GitHub設定、テスト方法などの詳細な手順が含まれています

## 🎯 サポートされるイベント

サーバーは以下のGitHubイベントを処理できます：

- **ping** - Webhook接続テスト
- **issues** - Issue作成、編集、クローズなど
- **issue_comment** - Issueコメント（コマンド検出をサポート）
- **push** - コードプッシュ
- その他のイベントもログに記録されます

## 🔧 APIエンドポイント

- `POST /webhook` - Webhookイベント受信エンドポイント
- `GET /health` - ヘルスチェックエンドポイント

## 🔒 セキュリティ機能

- HMAC-SHA256署名検証
- GitHub Webhook Secretのサポート
- 詳細なセキュリティログ記録
- 入力検証とエラー処理

## 📖 関連ドキュメント

- [Claude Code紹介](claude-code-intro.md) - AIプログラミングアシスタントツールの紹介
- [Webhookチュートリアル](webhook.md) - ゼロから始める完全なチュートリアル

## 🤝 貢献

このプロジェクトを改善するためのIssueやPull Requestの提出を歓迎します！

## 📄 ライセンス

このプロジェクトは学習およびデモンストレーション目的のみに使用されます。