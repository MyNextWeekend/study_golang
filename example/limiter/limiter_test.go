package limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	limiter := NewLimiter(2*time.Second, 2)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t.Format("2006-01-02 15:04:05"))
			if limiter.Allow() {
				fmt.Println("allow requests....")
			} else {
				fmt.Println("not allow request!!!!")
			}
		}
	}

}
