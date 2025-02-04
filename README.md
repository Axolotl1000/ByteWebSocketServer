# ByteWebSocketServer

![開源證書](https://img.shields.io/github/license/Axolotl1000/ByteWebSocketServer) ![Issue未解決數量](https://img.shields.io/github/issues/Axolotl1000/ByteWebSocketServer) ![PullRequest未解決數量](https://img.shields.io/github/issues-pr/Axolotl1000/ByteWebSocketServer) ![最近Commit時間](https://img.shields.io/github/last-commit/Axolotl1000/ByteWebSocketServer)

**簡單的 WebSocket 伺服器，提供身分驗證及房間內廣播功能**

---

## 執行前，你需要先準備

- MySQL 伺服器
- 公/私鑰 (一個用戶需要一對，使用者在不同房間不同步公鑰)，可以參考[#密鑰對生成](#密鑰對生成)，如果你已經有一對密鑰，可以參考[#驗證/轉換我的密鑰](#驗證/轉換我的密鑰)

## 密鑰對生成

**目前密鑰都使用`Ed25519`，請不要使用`RSA`或其他算法**

1. 前往 [https://cyphr.me/ed25519_tool/ed.html](https://cyphr.me/ed25519_tool/ed.html)
2. `Algorithm`選擇`Ed25519`
3. `Msg Encoding`及`Key Encoding`皆選擇`HEX`
4. 點擊`🎲🔑 Generate Random Key`生成全新的密鑰對
5. 私鑰，該網站稱為`Seed`，請**務必保存在安全的地方**，洩漏將導致**安全性為 0**
6. 公鑰，請加入伺服器資料庫，請參照[#可用腳本](#可用腳本)
   > 此操作不需要重啟伺服器

## 驗證/轉換我的密鑰

**目前密鑰都使用`Ed25519`，請不要使用`RSA`或其他算法**

1. 前往 [https://cyphr.me/ed25519_tool/ed.html](https://cyphr.me/ed25519_tool/ed.html)
2. `Algorithm`選擇`Ed25519`
3. `Msg Encoding`選擇`HEX`
4. `Key Encoding`選擇`你的密鑰編碼方式`
5. 在`Seed`填上你的私鑰。
6. 使用`⚙️ Public Key from Seed`按鈕生成公鑰或直接在`Public Key`填入你的公鑰
7. 點擊`🔏 Sign`驗證是否正確
8. 如果無法通過，請確認輸入是否正確或考慮[生成一對新的](#密鑰對生成)
9. 如果你的`Key Encoding`為`HEX`，則不需要轉換，如果為`Base64`，請使用選單將密鑰轉換為`HEX`格式
10. 網站會自動轉換為正確格式

## 可用腳本

`database_init.sql` - 初始化資料表，請先在目標資料庫後再執行  
`database_adduser.sql` - 添加使用者，請將指令的`roomId`、`userId`和`{HEXKEY}`替換後再執行  
`database_removeuser.sql` - 移除使用者，請將指令的`roomId`和`userId`替換後再執行

## 運行

1. 在[Release](https://github.com/Axolotl1000/ByteWebSocketServer/releases)頁面下載最新版本
   > 建議下載後驗證檔案的 SHA512 值是否與提供的相符
2. 設定環境變數
   > 這個軟體沒有設定檔案，請使用環境變數
   >
   > 需要設定下方的所有變數：
   >
   > - DB_USER = "資料庫使用者名稱"
   > - DB_PASSWORD = "資料庫密碼"
   > - DB_HOST = "資料庫位置"
   > - DB_NAME = "資料庫名稱"
   > - SERV_PORT = "伺服器監聽的連線埠"
3. 執行

## 編譯/單次運行

這個項目目前提供幾個腳本可以使用，請使用 sh 運行此腳本
`build.sh` - 編譯所有版本並生成 SHA512 值
`run.sh` - 重新編譯原始碼並執行
