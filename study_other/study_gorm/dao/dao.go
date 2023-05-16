package dao

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct { // 表名是users
	gorm.Model
	Name         string
	Email        string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
}
