package study_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitMySQL() {
	DB, _ = sql.Open("mysql", "study_go:123456789@tcp(106.55.186.222:3306)/study_go")
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println("connect mysql err:", err.Error())
		panic("connect mysql err:" + err.Error())
	}
	fmt.Println("connect mysql success")
}
