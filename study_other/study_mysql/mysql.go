package study_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitMySQL() {
	DB, _ = sql.Open("mysql", "study_golang:6xkiKTGzfPE7bXWE@tcp(106.55.186.222:3306)/study_golang")
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		panic("connect mysql err: " + err.Error())
	}
	fmt.Println("connect mysql success")
}
