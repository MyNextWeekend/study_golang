package limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	limiter := NewLimiter(2*time.Second, 2) // 每2秒生成1个令牌，桶容量为2

	for i := 1; i <= 10; i++ {
		if limiter.Allow() {
			fmt.Printf("Request %d passed\n", i)
		} else {
			fmt.Printf("Request %d blocked\n", i)
		}
		time.Sleep(500 * time.Millisecond) // 模拟请求间隔
	}

}
