@echo off

npm run build
go generate
go build -ldflags "-H windowsgui" -o lorca-vue.exe
