package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	mu sync.Mutex
	db *SQLManager
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleBinaryConnection(w http.ResponseWriter, r *http.Request) {

	roomID := r.URL.Query().Get("room")
	userID := r.URL.Query().Get("user")

	if roomID == "" || userID == "" {
		w.WriteHeader(400)
		w.Write([]byte("Empty room or user"))
		log.Printf("%s@%s 被中斷連線，原因：%s", userID, roomID, "嘗試連線的客戶端提供的訊息不正確")
		return
	}

	if connections[roomID] != nil {
		if _, exists := connections[roomID][userID]; exists {
			w.WriteHeader(400)
			w.Write([]byte("User already in the room"))
			log.Printf("%s@%s 被中斷連線，原因：%s", userID, roomID, "嘗試連線的客戶端訊息已經存在")
			return
		}
	}

	if !userExists(roomID, userID) {
		w.WriteHeader(400)
		w.Write([]byte("User Not Found Or Cannot Join This Room"))
		log.Printf("%s@%s 被中斷連線，原因：%s", userID, roomID, "嘗試連線的客戶端帳戶不存在")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket 升級失敗:", err)
		return
	}
	defer conn.Close()

	conn.WriteMessage(websocket.BinaryMessage, CONNECTED)

	code, e := createAuth(roomID, userID, conn)
	if e != "" {
		conn.WriteMessage(websocket.BinaryMessage, INTERNAL_ERROR)
		conn.Close()
		log.Printf("%s@%s 被中斷連線，原因：%s", userID, roomID, "無法創建驗證訊息\n"+e)
		return
	}

	log.Printf("用戶 %s 嘗試加入房間 %s，驗證中", userID, roomID)

	conn.WriteMessage(websocket.BinaryMessage, code)

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("用戶 %s 離線: %v", userID, err)
			break
		}

		if msgType != websocket.BinaryMessage {
			continue
		}

		if isAuth(roomID, userID) {
			if verify(roomID, userID, msg) {
				conn.WriteMessage(websocket.BinaryMessage, AUTH_SUCCESS)
				log.Printf("%s@%s %s", userID, roomID, "驗證通過")
				addConnection(roomID, userID, conn)
				log.Printf("用戶 %s 加入房間 %s", userID, roomID)
				removeAuth(roomID, userID)
				continue
			} else {
				conn.WriteMessage(websocket.BinaryMessage, AUTH_FAILED)
				conn.Close()
				log.Printf("%s@%s 被中斷連線，原因：%s", userID, roomID, "驗證不通過")
				removeAuth(roomID, userID)
				break
			}
		}

		if !isAuth(roomID, userID) {
			broadcastToRoom(roomID, userID, msg)
		}
	}

	removeConnection(roomID, userID)
	log.Printf("用戶 %s 離開房間 %s", userID, roomID)
}

func main() {
	var err error
	fmt.Println("正在連線至資料庫")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("SERV_PORT")

	if username == "" || password == "" || host == "" || dbname == "" || port == "" {
		log.Fatal("環境變數未設定，請確保 DB_USER, DB_PASSWORD, DB_HOST, DB_NAME SERV_PORT 都已設定")
		os.Exit(1)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, host, dbname)

	db, err = NewSQLManager(dsn)
	if err != nil {
		log.Fatal("伺服器啟動失敗:", err)
		os.Exit(1)
	}

	fmt.Println("成功連線至資料庫")

	fmt.Println("正在啟動伺服器")
	http.HandleFunc("/ws", handleBinaryConnection)
	fmt.Printf("Binary WebSocket 伺服器啟動於 :%s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("伺服器啟動失敗:", err)
		os.Exit(1)
	}
}
