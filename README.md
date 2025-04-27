# GoLang Windows Server Demo

## 构建

```batch
# windows
if (!(Test-Path "build")) { mkdir build }; Remove-Item "build/servicedemo.exe" -ErrorAction SilentlyContinue; go build -o build/servicedemo.exe main/main.go
```

# 使用

```batch
#  - 安装服务
build/servicedemo install
# - 卸载服务
build/servicedemo uninstall
# - 启动服务
build/servicedemo start
# - 停止服务
build/servicedemo stop
# - 重启服务
build/servicedemo restart
```