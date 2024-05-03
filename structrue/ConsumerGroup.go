package structrue

import (
	"sync"
	"time"
)

// 模拟 mq 消费组消费；

const (
	ChannelLength = 10
)

type MessageQueue struct {
	HearBeat time.Duration
	Mu       sync.Mutex
	cm       map[string]*ConsumerGroup
	mm       map[string][]Message
}

type Message struct {
	topic   string
	content string
}

func NewMQ(hb time.Duration) *MessageQueue {
	mq := &MessageQueue{}
	mq.mm = map[string][]Message{}
	mq.cm = map[string]*ConsumerGroup{}
	mq.Mu = sync.Mutex{}
	go mq.Push()
	return mq
}

func (mq *MessageQueue) Push() {
	for {
		time.Sleep(mq.HearBeat)
		mq.Mu.Lock()
		for topic, msgs := range mq.mm {
			if len(msgs) != 0 && mq.cm[topic] != nil {
				topush := msgs[0]
				mq.cm[topic].Broadcast(topush)
				mq.mm[topic] = mq.mm[topic][1:]
			}
		}
		mq.Mu.Unlock()
	}
}

func (mq *MessageQueue) Publish(topic string, content string) {
	NewMsg := Message{
		topic:   topic,
		content: content,
	}
	mq.Mu.Lock()
	mq.mm[topic] = append(mq.mm[topic], NewMsg)
	mq.Mu.Unlock()
}

type ConsumerGroup struct {
	mu        sync.Mutex
	consumers []*Consumer
}

func NewConsumerGroup() *ConsumerGroup {
	cg := &ConsumerGroup{}
	cg.consumers = []*Consumer{}
	cg.mu = sync.Mutex{}
	return cg
}

func (cg *ConsumerGroup) Broadcast(msg Message) {
	cg.mu.Lock()
	for _, c := range cg.consumers {
		select {
		case c.MsgCh <- msg:
		default:
		}
	}
	cg.mu.Unlock()
}

type Consumer struct {
	Topic string
	MsgCh chan Message
	mq    *MessageQueue
}

func NewConsumer(mq *MessageQueue) *Consumer {
	c := &Consumer{}
	c.mq = mq
	c.MsgCh = make(chan Message, ChannelLength)
	return c
}

func (c *Consumer) Subscribe(topic string) {
	c.Topic = topic
	if _, ok := c.mq.cm[topic]; !ok {
		c.mq.Mu.Lock()
		c.mq.cm[topic] = NewConsumerGroup()
		c.mq.Mu.Unlock()
	}
	c.mq.cm[topic].mu.Lock()
	c.mq.cm[topic].consumers = append(c.mq.cm[topic].consumers, c)
	c.mq.cm[topic].mu.Unlock()
}

func (c *Consumer) Consume() chan Message {
	return c.MsgCh
}
