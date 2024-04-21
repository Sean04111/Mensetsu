package util

import (
	"context"
	"testing"
)

func TestLimiter(t *testing.T) {
	limiter := NewLimiter(100, context.Background())
	for i := 0; i < 10000; i += 200 {
		limiter.Consume(200)
	}
}
