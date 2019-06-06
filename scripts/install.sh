set -e
echo 'Downloading dependencies...'
GO111MODULE=on go mod download
npm i -g -s serverless 