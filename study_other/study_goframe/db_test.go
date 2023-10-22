package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestInsert(t *testing.T) {
	// 获取默认配置的数据库对象(配置名称为"default")
	db := g.DB()
	var ctx = gctx.New()
	list1, err := db.GetAll(ctx, "select * from api limit 2")
	if err != nil {
		t.Error("db select err: " + err.Error())
	}
	t.Log("result: ", list1)
	list2, err := db.GetAll(ctx, "select * from user_info where age > ? and name like ?", g.Slice{18, "%john%"})
	if err != nil {
		t.Error("db select err: " + err.Error())
	}
	t.Log("result: ", list2)
	list3, err := db.GetAll(ctx, "select * from user_info where status=?", g.Slice{1})
	if err != nil {
		t.Error("db select err: " + err.Error())
	}
	t.Log("result: ", list3)
}
