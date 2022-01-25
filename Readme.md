# prototype

> 原型管理服务

## Setup

``` bash
# pack for linux
bee pack -be GOOS=linux -be GOARCH=amd64 -a prototype

# pack for macos
bee pack -be GOOS=darwin -be GOARCH=amd64 -a prototype

# pack for windows
bee pack -be GOOS=windows -be GOARCH=amd64 -a prototype

# create tmp folder
mkdir static/tmp
```

## swagger
```bash
# 下载swagger
bee run -downdoc=true

# 生成文档
# go install github.com/swaggo/swag/cmd/swag@latest
swag init -o ./swagger
```