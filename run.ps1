[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

Clear-Host

$env:GOOS = "windows"
$env:GOARCH = "amd64"
$env:DB_USER = "ByteWS"
$env:DB_PASSWORD = "ByteWS"
$env:DB_HOST = "127.0.0.1:3306"
$env:DB_NAME = "ByteWS"
$env:SERV_PORT = "80"
$executeable_filename = "bytews.exe"

Write-Host "✅ 環境變數已設定！"

$projectPath = "$PSScriptRoot"
Set-Location $projectPath

if (-Not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Host "❌ 未安裝 Go，請先安裝 Go ！" -ForegroundColor Red
    exit 1
}

Write-Host "📦 安裝 Go 依賴..."
go mod tidy

Write-Host "🚀 編譯 Go 應用程式..."
go build -o $executeable_filename

Write-Host "🏃 運行應用程式..."
./$executeable_filename