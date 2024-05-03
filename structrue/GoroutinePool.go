package structrue

import (
	"context"
	"sync/atomic"
)

//golang 协程池实现
//pool 主要目的在于控制 goroutine 数量
//考虑 data race

type Gpool struct {
	cap     uint32
	jobChan chan Job
	running atomic.Uint32
}

type Job struct {
	Work func()
	Ctx  context.Context
}

func NewPool(cap int) *Gpool {
	p := Gpool{}
	p.cap = uint32(cap)
	p.jobChan = make(chan Job, cap)
	go p.Worker()
	return &p
}

func (p *Gpool) Worker() {
	for {
		running := p.running.Load()
		if len(p.jobChan) > 0 && running < p.cap {
			if p.running.Load() != running {
				continue
			}

			task := <-p.jobChan

			p.running.Add(1)

			go func(ctx context.Context, f func()) {

				defer func() {
					load := p.running.Load()
					for !p.running.CompareAndSwap(load, load-1) {
						load = p.running.Load()
					}
				}()

				f()
			}(task.Ctx, task.Work)

		}
	}
}
func (p *Gpool) Do(f func()) {
	p.jobChan <- Job{Ctx: context.Background(), Work: f}
}
