# Notion Integration Setup

このプロジェクトでは、Notion APIを使用してページの一覧を表示する機能を追加しました。

## 必要な環境変数

バックエンドの環境変数を設定してください：

```bash
# Notion API Token
NOTION_TOKEN=your_notion_integration_token

# Notion Database ID
NOTION_DATABASE_ID=your_database_id
```

## Notion API の設定手順

### 1. Notion Integration の作成

1. [Notion Developers](https://www.notion.so/my-integrations) にアクセス
2. "New integration" をクリック
3. 統合名を入力（例：My Blog Integration）
4. "Submit" をクリック
5. 生成された "Internal Integration Token" をコピー

### 2. データベースの共有設定

1. Notionで対象のデータベースを開く
2. 右上の "Share" ボタンをクリック
3. "Invite" セクションで作成した統合を検索
4. 統合を選択して "Invite" をクリック

### 3. データベースIDの取得

1. データベースのURLをコピー
2. URLの形式：`https://www.notion.so/workspace/DATABASE_ID?v=...`
3. `DATABASE_ID` の部分をコピー

## 環境変数の設定

### Docker Compose を使用する場合

`docker-compose.yml` ファイルに環境変数を追加：

```yaml
services:
  backend:
    environment:
      - NOTION_TOKEN=your_notion_integration_token
      - NOTION_DATABASE_ID=your_database_id
```

### ローカル開発の場合

`.env` ファイルを作成：

```env
NOTION_TOKEN=your_notion_integration_token
NOTION_DATABASE_ID=your_database_id
```

## 使用方法

1. 環境変数を設定
2. アプリケーションを起動
3. ナビゲーションバーの "Notion" リンクをクリック
4. Notionページの一覧が表示されます

## 機能

- Notionデータベースからページ一覧を取得
- ページタイトル、作成日時、更新日時を表示
- ページアイコン（絵文字）を表示
- Notionでページを開く機能
- レスポンシブデザイン
- エラーハンドリング

## API エンドポイント

- `GET /notion/pages` - Notionページ一覧を取得

## トラブルシューティング

### よくある問題

1. **"Notion API not configured" エラー**
   - `NOTION_TOKEN` 環境変数が設定されているか確認

2. **"Notion database ID not configured" エラー**
   - `NOTION_DATABASE_ID` 環境変数が設定されているか確認

3. **"Failed to fetch pages from Notion" エラー**
   - Notion統合がデータベースにアクセス権限を持っているか確認
   - データベースIDが正しいか確認

### デバッグ

バックエンドのログを確認して、詳細なエラーメッセージを確認してください。