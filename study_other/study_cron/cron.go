package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func main() {
	c := cron.New()
	err := c.AddFunc("*/5 * * * * *", func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "每5秒执行一次")
	})
	if err != nil {
		fmt.Println("addFunc err:", err.Error())
		return
	}
	c.Start()
	// 阻塞
	select {}
}
