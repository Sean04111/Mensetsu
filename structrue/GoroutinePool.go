package structrue

import (
	"context"
)

//golang 协程池实现

type Gpool struct {
	cap     uint8
	jobChan chan Job
	running uint8
}

type Job struct {
	Work func()
	Ctx  context.Context
}

func NewPool(cap int) *Gpool {
	p := Gpool{}
	p.cap = uint8(cap)
	p.jobChan = make(chan Job, 10)
	go p.Worker()
	return &p
}

func (p *Gpool) Worker() {
	for {
		if len(p.jobChan) > 0 && p.running < p.cap {
			task := <-p.jobChan
			go func(ctx context.Context, f func()) {
				p.running++
				defer func ()  {
					p.running--
				}()
				f()
				return
			}(task.Ctx, task.Work)
		}
	}
}
func (p *Gpool) Do(f func()) {
	p.jobChan<-Job{Ctx: context.Background(),Work: f}
}
