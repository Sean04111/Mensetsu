package structrue

import (
	"fmt"
	"testing"
	"time"
)

func TestMQ(m *testing.T) {
	fmt.Println("Start")
	mq := NewMQ(50 * time.Microsecond)

	for i := 0; i < 5; i++ {
		num := i
		go func() {
			c := NewConsumer(mq)
			if num%2 == 0 {
				c.Subscribe("xbc")
			} else {
				c.Subscribe("zh")
			}
			for {
				select {
				case msg := <-c.Consume():
					fmt.Println("consumer ", num, "consumed : ", msg.content)
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second)

	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				mq.Publish("xbc", "xbc")
			} else {
				mq.Publish("zh", "zh")
			}
			time.Sleep(time.Second)
		}
	}()

	select {}
}
