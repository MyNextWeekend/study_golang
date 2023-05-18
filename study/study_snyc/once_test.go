package study_snyc

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func mothed(index int) {
	time.Sleep(1 * time.Second)
	fmt.Println("hello world index:", index)
}

func TestName(t *testing.T) {
	once := &sync.Once{}
	for i := 0; i < 100; i++ {

		go func(idx int) {
			once.Do(func() {
				mothed(idx)
			})
		}(i)
	}

	time.Sleep(10 * time.Second)
}
