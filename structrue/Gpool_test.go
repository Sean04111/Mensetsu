package structrue

import (
	"fmt"
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

func TestChannelRange(t *testing.T) {
	ch := make(chan int,2)
	go func() {
			ch <- 1
	}()
	time.Sleep(time.Second)

	fmt.Println("a")
	fmt.Println(len(ch))
	for v := range ch {
		fmt.Println("got :", v)
	}
}
