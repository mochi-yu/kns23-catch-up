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
- APIサーバポート番号：8085
- フロントエンドポート番号：
