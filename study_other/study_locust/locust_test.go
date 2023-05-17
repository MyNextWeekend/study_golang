package study_locust

import (
	"github.com/myzhan/boomer"
	"log"
	"os"
	"testing"
	"time"
)

func foo() {
	time.Sleep(5 * time.Second)
	start := time.Now()

	time.Sleep(3 * time.Second)
	log.Println("触发了")

	elapsed := time.Since(start)
	boomer.RecordSuccess("http", "foo", elapsed.Nanoseconds()/int64(time.Millisecond), int64(10))
}

func TestLocust(t *testing.T) {
	logFile, err := os.OpenFile("./test_locust/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)

	task1 := &boomer.Task{
		Name:   "foo",
		Weight: 10,
		Fn:     foo,
	}

	boomer.Run(task1)
}
