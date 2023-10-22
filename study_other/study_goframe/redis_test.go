package main

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestGet(t *testing.T) {
	var ctx = gctx.New()
	value, err := g.Redis().Get(ctx, "key")
	if err != nil {
		t.Error("redis get err: " + err.Error())
	}
	t.Log("result: ", value)
}

func TestSet(t *testing.T) {
	var ctx = gctx.New()
	value, err := g.Redis().Set(ctx, "key", "zhangsan")
	if err != nil {
		t.Error("redis get err: " + err.Error())
	}
	t.Log("result: ", value)
}
