package main

import (
	"github.com/myzhan/boomer"
	"time"
)

func foo() {
	start := time.Now()
	time.Sleep(3 * time.Second)
	elapsed := time.Since(start)

	boomer.RecordSuccess("http", "foo", elapsed.Nanoseconds()/int64(time.Millisecond), int64(10))
}

func main() {
	task1 := &boomer.Task{
		Name:   "foo",
		Weight: 10,
		Fn:     foo,
	}

	boomer.Run(task1)
}
