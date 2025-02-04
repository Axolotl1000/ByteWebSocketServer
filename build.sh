#!/bin/bash

ScriptDir=$(dirname "$(realpath "$0")")
cd "$ScriptDir"

buildDir="$ScriptDir/build"
if [ ! -d "$buildDir" ]; then
    mkdir -p "$buildDir"
fi

sha512File="$buildDir/SHA512.md"
if [ -f "$sha512File" ]; then
    rm "$sha512File"
fi

targets=(
    "windows amd64 .exe"
    "windows arm64 .exe"
    "windows 386 .exe"
    "windows arm .exe"
    "linux amd64"
    "linux arm64"
    "linux 386"
    "linux arm"
    "darwin amd64"
    "darwin arm64"
)

for target in "${targets[@]}"; do
    IFS=' ' read -r GOOS GOARCH Ext <<< "$target"
    echo "🛠️ 編譯 $GOOS-$GOARCH..."

    export GOOS
    export GOARCH

    filename="app_${GOOS}_${GOARCH}${Ext}"

    # 編譯程式
    go build -o "$buildDir/$filename"

    if [ $? -eq 0 ]; then
        echo "✅ 編譯成功：$buildDir/$filename"

        # 計算 SHA512
        sha512Value=$(sha512sum "$buildDir/$filename" | awk '{ print $1 }')
        echo "$filename : SHA512::$sha512Value"

        # 寫入 SHA512.md
        echo "- $filename: \`\`$sha512Value\`\`" >> "$sha512File"
    else
        echo "❌ 編譯失敗：$buildDir/$filename"
    fi
done

echo "🎉 所有編譯完成。"
echo "📄 SHA512 檔案已生成：$sha512File"
