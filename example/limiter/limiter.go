package limiter

import (
	"sync"
	"time"
)

type TokenBuket struct {
	mu                sync.Mutex
	size              int           //桶容量
	count             int           //当前令牌数量
	rateLimit         time.Duration //多久填充一个令牌
	lastFillTokenTime time.Time     //最后一次成功填充的时间
}

// 填充令牌数量
func (t *TokenBuket) fillToken() {
	if num := t.getFillTokenCount(); num > 0 {
		// 只有真实执行填充之后，才会修改时间
		t.lastFillTokenTime = time.Now()
		t.count += num
	}
}

// 返回需要填充令牌的数量
func (t *TokenBuket) getFillTokenCount() int {
	if t.count >= t.size {
		return 0
	}
	now := time.Now()
	count := int(now.Sub(t.lastFillTokenTime) / t.rateLimit)
	if t.size-t.count >= count {
		return count
	}
	return t.size - t.count
}

// 判断令牌数量是否足够
func (t *TokenBuket) allow() bool {
	// 先填充令牌
	t.fillToken()
	// 查看令牌数量是否足够
	if t.count > 0 {
		t.count--
		return true
	}
	return false
}

type Limiter struct {
	tb *TokenBuket
}

func NewLimiter(r time.Duration, size int) *Limiter {
	return &Limiter{
		tb: &TokenBuket{
			size:      size,
			rateLimit: r,
			count:     size,
		},
	}
}

func (l *Limiter) Allow() bool {
	l.tb.mu.Lock()
	defer l.tb.mu.Unlock()

	return l.tb.allow()
}
