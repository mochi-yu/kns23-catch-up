# kns23-catch-up
高校の同窓会に向けて、各自の近況を報告できるSNSです。

## 開発方法
### 開発手順
#### 基本の手順
1. 実装内容のissueを作成する
2. `feat/#{issue番号}`として作業用ブランチを作成する
3. コミットメッセージに`#{issue番号}`を含めてコミットする
4. PRを作成する
5. レビューを受けてmainにマージする

#### フロントエンド開発時の手順
1. `cd frontend`で移動する
2. 初回時は`make setup`を実行する
3. 開発用サーバを`make dev`として起動する
4. バックエンドとの接続をする際は、以下の手順を参照する

#### api開発時の手順
1. 初回時は`make build`を実行する
2. `make up`として開発用サーバを起動する
3. goのファイルを編集して保存すると、`webapi-kns23`のコンテナはホットリロードされる
4. apiの編集に関しては [/api/README.md](api/README.md)を参照

#### API仕様の変更時
1. `document/api/src`以下のyamlファイルを編集する
2. `make lint`を実行し、エラーが出ていないことを確認する
3. コマンドが見つからなかったら`make setup`を実行する
4. `make compile`を実行し`document/api/api.yaml`を生成する

### リリース手順
1. releaseブランチにマージする

## 基本設計
### UI設計
- Figmaにて作成

### システム設計
![システム構成:2024/01/09](./document/images/system-architecture.jpg)

### DB設計
- 参照: [DB.md](./document/DB.md)

### API設計
- 参照: [api.yaml](./document/api/api.yaml)

## 基本情報
### ローカル実行時の情報
- APIサーバポート番号：8086
- フロントエンドポート番号：3085
