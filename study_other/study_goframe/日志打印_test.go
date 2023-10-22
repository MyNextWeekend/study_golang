package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestLog(t *testing.T) {
	ctx := context.Background()
	g.Log().Debug(ctx, "这个是Debug日志。。。")
	g.Log().Info(ctx, "这个是Info日志。。。")
	g.Log().Warning(ctx, "这个是Warning日志。。。")
	g.Log().Error(ctx, "这个是Error日志。。。")
	g.Log().Error(ctx, "这个是Error日志。。。")
}
