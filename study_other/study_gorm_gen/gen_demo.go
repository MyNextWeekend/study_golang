package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./study_other/study_gorm_gen/dal",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery, // generate mode
	})

	dsn := "study_golang:6xkiKTGzfPE7bXWE@(106.55.186.222:3306)/study_golang?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("open db err: " + err.Error())
	}
	g.UseDB(db) // 设置目标数据库

	// 生成所有表
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
