package study

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("超过三秒钟了呀")
	default:
		fmt.Print("999")

	}
}
