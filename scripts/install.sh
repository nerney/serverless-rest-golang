set -e
echo 'Downloading dependencies...'
npm i -g -s serverless 
GO111MODULE=on go mod download