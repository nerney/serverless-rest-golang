set -e 
echo "Building..."
rm -rf ./bin
GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/rest lambda/main.go
