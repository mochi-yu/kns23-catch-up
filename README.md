# kns23-catch-up
高校の同窓会に向けて、各自の近況を報告できるSNSです。

## 基本設計
### UI設計
- Figmaにて作成

### システム設計
![システム構成:2024/01/09](./document/images/system-architecture.jpg)

### DB設計
- 参照: [DB.md](./document/DB.md)

### API設計
- 参照: [api.yaml](./document/api/api.yaml)

## 開発方法
### 開発手順
1. 実装内容のissueを作成する
2. `feat/#{issue番号}`として作業用ブランチを作成する
3. コミットメッセージに`#{issue番号}`を含めてコミットする
4. PRを作成する
5. レビューを受けてmainにマージする

### リリース手順
1. releaseブランチにマージする