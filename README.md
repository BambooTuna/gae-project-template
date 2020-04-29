## ローカル動作確認
```bash
cd ./terraform
ssh-keygen -t rsa -f my-ssh-key -C ${SSH_USERNAME}
```
### 環境構築
```bash
chmod +x ${PWD}

export HOME_DIE=${PWD}

export TF_VAR_GOOGLE_PROJECT_ID=develop-275713
export TF_VAR_GOOGLE_COMPUTE_REGION=asia-northeast1
export TF_VAR_GOOGLE_COMPUTE_ZONE=asia-northeast1-a
export TF_VAR_GOOGLE_CREDENTIALS_JSON_PATH=./account.json
export TF_VAR_SSH_PORT=60001
export SSH_USERNAME=bambootuna

cd ${HOME_DIE}/terraform

export TF_VAR_SSH_PUB_KEY=${SSH_USERNAME}:$(cat my-ssh-key.pub)
sed -e "s/<%SHH_PORT%>/${TF_VAR_SSH_PORT}/" cloud-config.tpl > cloud-config
terraform init
sh import.sh
sh apply.sh

```

### GAEデプロイ
```bash

npm install --prefix ${HOME_DIE}/front
npm run build --prefix ${HOME_DIE}/front
rm -rf ${HOME_DIE}/apiServer/dist
mv ${HOME_DIE}/front/dist ${HOME_DIE}/apiServer/dist

export VPC_ACCESS_CONNECTOR_NAME=projects/${TF_VAR_GOOGLE_PROJECT_ID}/locations/${TF_VAR_GOOGLE_COMPUTE_REGION}/connectors/access-connector
export API_SERVER_ENDPOINT=http://localhost

cd ${HOME_DIE}/apiServer
go get -v -t -d ./...
cat ./app.tpl.yml | ./extcat.sh > ./app.yml

echo 'github-actions@${TF_VAR_GOOGLE_PROJECT_ID}.iam.gserviceaccount.com' | gcloud auth activate-service-account --key-file ${HOME_DIE}/terraform/account.json
gcloud app deploy app.yml --project ${TF_VAR_GOOGLE_PROJECT_ID} --quiet
```

### サーバーデプロイ
```bash
if cd middleware; then git pull; else git clone https://github.com/BambooTuna/middleware.git middleware; fi

```
