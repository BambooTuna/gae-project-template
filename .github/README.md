## 設定すべき環境変数

### GCP関連
- GOOGLE_PROJECT_ID
- GOOGLE_COMPUTE_REGION

    Optional
    default: asia-northeast1

- GOOGLE_COMPUTE_ZONE

    Optional
    default: asia-northeast1-a

- GOOGLE_SERVICE_KEY

    サービスアカンウトをbase64エンコードした文字列
    `base64 -i [.json file path]`

### SSH関連

1. terraformでGCEサーバーを作成する際に使用する
2. ssh経由でGCEサーバーにデプロイする際に使用する

- SSH_USERNAME
- SSH_KEY

    ssh秘密鍵をbase64エンコードした文字列
    `base64 -i [ssh key file path]`

- SSH_KEY_PUB

    ssh公開鍵をbase64エンコードした文字列
    `base64 -i [ssh key pub file path]`

- SSH_HOST

    ssh接続するGCEサーバーの外部IP

- SSH_PORT

    ssh接続のポート
    空いていればなんでもよい。基本は22だがセキュリティー上変更した方がいい
