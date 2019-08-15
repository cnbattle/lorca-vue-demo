#!/bin/sh

if [ ! -d "./dist/" ];then
  npm run build
fi

go generate

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui" -o lorca-vue.exe
