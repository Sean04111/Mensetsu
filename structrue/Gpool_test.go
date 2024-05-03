package structrue

import (
	"testing"
	"time"
)

func TestGpool(t *testing.T) {
	pool := NewPool(10)
	for i := 0; i < 20; i++ {
		pool.Do(func() {
			for j := 0; j < 10; j++ {
				time.Sleep(1 * time.Second)
			}
		})
	}
	select {}
}
