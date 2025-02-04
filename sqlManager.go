package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type SQLManager struct {
	db *sql.DB
}

func NewSQLManager(dsn string) (*SQLManager, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("無法建立資料庫連接: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("無法 ping 資料庫: %v", err)
	}

	return &SQLManager{db: db}, nil
}

func (s *SQLManager) ExecPreparedStatement(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("準備 SQL 語句失敗: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, fmt.Errorf("執行 SQL 語句失敗: %v", err)
	}
	return result, nil
}

func (s *SQLManager) QueryPreparedStatement(query string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("準備 SQL 查詢失敗: %v", err)
	}

	rows, err := stmt.Query(args...)
	if err != nil {
		stmt.Close() // 若發生錯誤，需手動關閉 stmt
		return nil, fmt.Errorf("查詢資料時發生錯誤: %v", err)
	}
	return rows, nil
}

func (s *SQLManager) QueryRowPreparedStatement(query string, args ...interface{}) *sql.Row {
	// 直接使用 db.QueryRow，避免 stmt.Close() 問題
	return s.db.QueryRow(query, args...)
}

func (s *SQLManager) Close() {
	err := s.db.Close()
	if err != nil {
		log.Printf("關閉資料庫連接時發生錯誤: %v", err)
	}
}
