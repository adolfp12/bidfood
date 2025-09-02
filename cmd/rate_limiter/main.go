package main

import (
	"fmt"
	"time"
)

func main() {
	limiter := NewRateLimiter(5) // 5 operations per second

	for i := 0; i < 10; i++ {
		limiter.Allow()
		fmt.Println("Operation", i, time.Now())
	}
}

type Limiter struct {
	ConcurrentProcess int
	DoProcess         chan int
}

func NewRateLimiter(n int) *Limiter {
	return &Limiter{
		ConcurrentProcess: n,
		DoProcess:         make(chan int),
	}
}

func (l *Limiter) Allow() {
	go l.Run(l.DoProcess, 1) // running process
	time.Sleep(time.Second / time.Duration(l.ConcurrentProcess))
	<-l.DoProcess
}

func (l *Limiter) Run(tickChan chan<- int, processNum int) {
	tickChan <- processNum
}
