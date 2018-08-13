package main

import (
	"time"
	"log"
)

// Intervaler 是个接口用来让调用者自定义poller轮询时间间隔
type Intervaler interface {
	Interval() time.Duration
}

// IntervalerFunc 用来将 func() time.Duration 转化成 Intervaler
type IntervalerFunc func() time.Duration

func (intervalerFunc IntervalerFunc) Interval() time.Duration {
	return intervalerFunc()
}

type Poller struct {
	//要执行的方法
	do           func() error
	//用于调用者传递停止信号
	cancel chan int
	//下次调用的时间间隔
	nextInterval Intervaler
}

// Poll 轮询
func (poller *Poller) Poll() {
	for {
		select {
		case <-poller.cancel:
			return
		case <-time.After(poller.nextInterval.Interval()):
			go func() {
				if err := poller.do(); err != nil {
					log.Printf("Poll poller.go: polling method returns a error: %v", err)
					// 或者结束整个轮询
					// poller.Cancel()
				}
			}()
		}
		println("aaa")
	}
}

// Cancel 向 cancel 发送信号
func (poller *Poller) Cancel() {
	println("Polling stopped")
	poller.cancel <- 1
}

// NewPoller 创建一个新的 Poller
func NewPoller(intervaler Intervaler, do func() error) *Poller {
	return &Poller{do: do, cancel: make(chan int), nextInterval: intervaler}
}

func main() {
	// 自定义 Intervaler
	base := time.Second * 0
	interval := IntervalerFunc(func() time.Duration {
		next := base
		base += 500 * time.Millisecond
		return next
	})

	// 创建一个 poller
	poller := NewPoller(interval,
		func() error {
			// 4秒后 输出 ping!
			time.AfterFunc(time.Second*4, func() { println("ping!")
				println(time.Now().String())
			})
			return nil
		})
	// 10 秒后停止 polling
	time.AfterFunc(time.Second*10, poller.Cancel)
	// 开始 polling
	println(time.Now().String())
	poller.Poll()
}