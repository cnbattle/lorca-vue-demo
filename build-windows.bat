@echo off

if !exist ".\dist\" (
    npm run build
)

go generate

go build -ldflags "-H windowsgui" -o lorca-vue.exe
