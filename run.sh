#!/bin/bash

export GOOS=linux
export GOARCH=amd64
export DB_USER="ByteWS"
export DB_PASSWORD="ByteWS"
export DB_HOST="127.0.0.1:3306"
export DB_NAME="ByteWS"
export SERV_PORT="80"
executeable_filename="bytews.exe"

echo "✅ 環境變數已設定！"

projectPath=$(dirname "$(realpath "$0")")
cd "$projectPath"

if ! command -v go &> /dev/null; then
    echo "❌ 未安裝 Go，請先安裝 Go ！" >&2
    exit 1
fi

echo "📦 安裝 Go 依賴..."
go mod tidy

echo "🚀 編譯 Go 應用程式..."
go build -buildvcs=false -o "$executeable_filename"

echo "🏃 運行應用程式..."
./"$executeable_filename"
