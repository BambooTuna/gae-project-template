# gae-project-template
フロントとREST APIのリバースプロキシサーバーをGAEに乗せ、APIサーバー・DBをGCEに乗せて、
間の通信を`vpc access connector`で行うプロジェクトのテンプレート

- Pros
1. GAE・GCEの無料枠に強引に収めることができる
2. GAEのドメインはhttpsなので、混合コンテンツのデフォルトブロックを防ぐことができる
※ロードバランサーなどを使ってhttps化するのは面倒

- Cons
1. https化する必要がないのであれば、普通にGCE+docker-composeを使った方が楽
2. スケールさせたい場合は素直にk8sを使った方がいい

## 準備
1. ssh経由デプロイ用の鍵作成

`$ ssh-keygen -t rsa -f my-ssh-key -C [任意のsshユーザーネーム]`

2. CI/CDに必要な環境変数をGithubActionsに設定する

[設定すべき環境変数一覧](#設定すべき環境変数)

## ワークフロー
- master
ビルドとデプロイ

- staging
terraformによるGCP環境構築・更新

## 設定すべき環境変数

### GCP関連
- GOOGLE_PROJECT_ID
- GOOGLE_COMPUTE_REGION

    例: asia-northeast1

- GOOGLE_COMPUTE_ZONE

    例: asia-northeast1-a

- GOOGLE_SERVICE_KEY

    サービスアカンウトをbase64エンコードした文字列
    `base64 -i [.json file path]`

### SSH関連

1. terraformでGCEサーバーを作成する際に使用する
2. ssh経由でGCEサーバーにデプロイする際に使用する

- SSH_USERNAME

    sshユーザーネーム

- SSH_KEY

    ssh秘密鍵をbase64エンコードした文字列
    `base64 -i my-ssh-key`

- SSH_KEY_PUB

    ssh公開鍵をbase64エンコードした文字列
    `base64 -i my-ssh-key.pub`

- SSH_HOST

    ssh接続するGCEサーバーの外部IP

- SSH_PORT

    ssh接続のポート
    空いていればなんでもよい。基本は22だがセキュリティー上変更した方がいい


## ローカル動作確認

### 準備
1. Run

```bash
$ chmod +x local-run.sh
$ ./local-run.sh
```

2. Curl
```bash
$ curl -X POST -d '{"text":"sample message"}' http://localhost:8080/v1/sample
```

3. swagger
[swagger page](http://localhost:8080/v1/swagger/index.html)

