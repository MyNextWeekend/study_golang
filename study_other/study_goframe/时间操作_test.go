package main

import (
	"github.com/gogf/gf/v2/os/gtime"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	// 通过时间字符串创建
	gtime.New("2020-10-24 12:00:00")
	// 通过time.Time对象创建
	gtime.New(time.Now())
	// 通过时间戳(秒)创建
	gtime.New(1603710586)
	// 通过时间戳(纳秒)创建
	gtime.New(1603710586660409000)
}
