$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition
Set-Location $ScriptDir

$buildDir = "$ScriptDir\build"
if (-not (Test-Path $buildDir)) {
    New-Item -ItemType Directory -Path $buildDir
}

$sha512File = "$buildDir\SHA512.md"
if (Test-Path $sha512File) {
    Remove-Item $sha512File
}

$targets = @(
    @{ GOOS = "windows"; GOARCH = "amd64"; Ext = ".exe" },
    @{ GOOS = "windows"; GOARCH = "arm64"; Ext = ".exe" },
    @{ GOOS = "windows"; GOARCH = "386"; Ext = ".exe" },
    @{ GOOS = "windows"; GOARCH = "arm"; Ext = ".exe" },
    @{ GOOS = "linux"; GOARCH = "amd64"; Ext = "" },
    @{ GOOS = "linux"; GOARCH = "arm64"; Ext = "" },
    @{ GOOS = "linux"; GOARCH = "386"; Ext = "" },
    @{ GOOS = "linux"; GOARCH = "arm"; Ext = "" },
    @{ GOOS = "darwin"; GOARCH = "amd64"; Ext = "" },
    @{ GOOS = "darwin"; GOARCH = "arm64"; Ext = "" }
)

foreach ($target in $targets) {
    Write-Host "🛠️ 編譯 $($target.GOOS)-$($target.GOARCH)..." -ForegroundColor Cyan
    
    $env:GOOS = $target.GOOS
    $env:GOARCH = $target.GOARCH
    
    # 設定輸出檔案名稱和副檔名
    $filename = "app_$env:GOOS" + "_" + "$env:GOARCH" + $target.Ext

    go build -o "$buildDir\$filename"
    
    if ($?) {
        Write-Host "✅ 編譯成功：$buildDir\$filename" -ForegroundColor Green
        
        # 生成SHA512 hash
        $sha512 = Get-FileHash "$buildDir\$filename" -Algorithm SHA512
        $sha512Value = $sha512.Hash
        Write-Host "$filename : SHA512::$sha512Value" -ForegroundColor Yellow
        
        # 添加進SHA512.md檔案
        $mdEntry = "- $filename" + ": ``$sha512Value``"
        Add-Content -Path $sha512File -Value $mdEntry
    }
    else {
        Write-Host "❌ 編譯失敗：$buildDir\$filename" -ForegroundColor Red
    }
}

Write-Host "🎉 所有編譯完成。" -ForegroundColor Magenta
Write-Host "📄 SHA512 檔案已生成：$sha512File" -ForegroundColor Blue
