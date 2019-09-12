package model

import (
	"fmt"
	"log"
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {  // 对数据库进行封装
	Self *sql.DB
}

var DB *Database

func getDatabase(username, password, host, port, dbname string) (*sql.DB, error) {
    address := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
    fmt.Println("Start to connect:", address)
    db, err := sql.Open("mysql", address)
    if err != nil {
    	return nil, err
    }
    err = db.Ping()
    if err != nil {
    	return nil, err
    }
    return db, nil
}

func (db *Database) Init() {
	newdb, err := getDatabase("root", "ginny", "127.0.0.1", "3306", "go_blog")
	if err != nil {
		log.Fatal(err)  // 如果有错误，那么打印，而不是直接退出
	}
	DB = &Database{
		Self: newdb,
	}
}

func (db *Database) Close() {
	DB.Self.Close()
}
