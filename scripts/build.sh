rm -rf ./bin
export GO111MODULE=on
export GOOS=linux
go build -ldflags="-s -w" -o bin/rest lambdas/rest/main.go
