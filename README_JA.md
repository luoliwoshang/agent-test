# GitHub Webhook デモプロジェクト

Go サーバー実装と詳細なチュートリアルドキュメントを含む、GitHub Webhook 機能の学習とデモンストレーションのための完全なプロジェクトです。

## 🚀 プロジェクト概要

このプロジェクトは、以下をサポートする完全な GitHub Webhook 受信サーバーを提供します：

- 様々な GitHub イベントの受信と処理（Issues、Comments、Push など）
- セキュリティのための HMAC-SHA256 署名検証
- ローカル開発と Render プラットフォームデプロイメントのサポート
- 詳細なログ出力とイベント処理
- ヘルスチェックエンドポイント

## 📁 プロジェクト構造

```
├── README.md              # プロジェクトドキュメント（中国語）
├── README_EN.md           # プロジェクトドキュメント（英語）
├── README_FR.md           # プロジェクトドキュメント（フランス語）
├── README_JA.md           # プロジェクトドキュメント（日本語）
├── webhook.md             # 完全な Webhook チュートリアル
├── claude-code-intro.md   # Claude Code ツール紹介
├── go.mod                 # Go モジュール設定（ルート）
├── render.yaml            # Render プラットフォームデプロイ設定
└── webhook-demo/          # Webhook サーバー実装
    ├── go.mod            # Go モジュール設定
    └── server.go         # メインサーバーコード
```

## 🛠️ クイックスタート

### 前提条件

- Go 1.21 以上
- Git

### ローカル開発

1. **リポジトリのクローン**
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

### Webhook Secret の設定（推奨）

1. **シークレットの生成**
   ```bash
   openssl rand -hex 20
   ```

2. **環境変数の設定**
   ```bash
   export WEBHOOK_SECRET=your-generated-secret
   cd webhook-demo
   go run server.go
   ```

## 🌐 Render へのデプロイ

プロジェクトには `render.yaml` 設定ファイルが含まれており、Render プラットフォームへのワンクリックデプロイが可能です：

1. GitHub リポジトリを Render に接続
2. Web Service を選択
3. Render が自動的に `render.yaml` の設定を使用

デプロイ後、自動的に以下が設定されます：
- `WEBHOOK_SECRET` 環境変数の設定
- ヘルスチェックの設定
- 自動デプロイの有効化

## 📚 使用方法チュートリアル

詳細な使用方法チュートリアルについては、以下をご覧ください：
- [**webhook.md**](webhook.md) - 完全な GitHub Webhook 設定と使用方法チュートリアル
- トンネリング、GitHub 設定、テスト方法、詳細な手順を含みます

## 🎯 サポートされるイベント

サーバーは以下の GitHub イベントを処理できます：

- **ping** - Webhook 接続テスト
- **issues** - Issue の作成、編集、クローズなど
- **issue_comment** - Issue コメント（コマンド検出をサポート）
- **push** - コードプッシュ
- その他のイベントもログに記録されます

## 🔧 API エンドポイント

- `POST /webhook` - Webhook イベント受信エンドポイント
- `GET /health` - ヘルスチェックエンドポイント

## 🔒 セキュリティ機能

- HMAC-SHA256 署名検証
- GitHub Webhook Secret サポート
- 詳細なセキュリティログ記録
- 入力検証とエラーハンドリング

## 📖 関連ドキュメント

- [Claude Code 紹介](claude-code-intro.md) - AI プログラミングアシスタントツールの紹介
- [Webhook チュートリアル](webhook.md) - ゼロから始める完全なチュートリアル

## 🤝 貢献

このプロジェクトを改善するための Issue や Pull Request を歓迎します！

## 📄 ライセンス

このプロジェクトは学習とデモンストレーション目的のみです。