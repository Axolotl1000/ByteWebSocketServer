package main

import (
	"encoding/hex"
	"fmt"

	"github.com/gorilla/websocket"
)

type Auth struct {
	connection *websocket.Conn
	authCheck  []byte
}

var auth = make(map[string]map[string]Auth)

func createAuth(room, user string, conn *websocket.Conn) ([]byte, string) {
	if auth[room] == nil {
		auth[room] = make(map[string]Auth)
	}
	key, err := Generate(32)
	if err != nil {
		return nil, err.Error()
	}
	key = append(AUTH, key...)
	auth[room][user] = Auth{
		connection: conn,
		authCheck:  key,
	}

	return key, ""
}

func removeAuth(room, user string) {
	if auth[room] == nil {
		return
	}

	delete(auth[room], user)

	if len(auth[room]) == 0 {
		delete(auth, room)
	}
}

func isAuth(room, user string) bool {
	if auth[room] == nil {
		return false
	}
	if _, exists := auth[room][user]; !exists {
		return false
	}

	return true
}

func userExists(room, user string) bool {
	const search string = "SELECT `user` FROM `public_keys` WHERE `room` = ? AND `user` = ?;"
	res, err := db.QueryPreparedStatement(search, room, user)
	if err != nil {
		fmt.Printf("查詢失敗：%s\n", err)
		return false
	}
	defer res.Close()

	return res.Next()
}

func verify(room, user string, msg []byte) bool {
	if !isAuth(room, user) {
		return false
	}

	const search string = "SELECT `key` FROM `public_keys` WHERE `room` = ? AND `user` = ?;"
	res, err := db.QueryPreparedStatement(search, room, user)
	if err != nil {
		fmt.Printf("查詢失敗：%s\n", err)
		return false
	}
	defer res.Close()

	if !res.Next() {
		return false
	}

	var keyHex string
	if err := res.Scan(&keyHex); err != nil {
		fmt.Printf("讀取查詢結果失敗：%s\n", err)
		return false
	}

	p_key, err := hex.DecodeString(keyHex)
	if err != nil {
		fmt.Printf("公鑰 HEX 解析失敗：%s\n", err)
		return false
	}

	return Verify(p_key, auth[room][user].authCheck, msg)
}
