package Week06

import (
	"errors"
	"sync"
	"time"
)

type Limiter struct {
	lock      sync.Mutex
	once      sync.Once
	totalTime time.Duration //窗口所持续的时间
	tick      time.Duration //每一格所占的时间
	window    []int64
	index     int   //窗口中所处当前位置
	max       int64 //最大计数

	stop chan struct{}
}

func NewLimiter(totalTime, tick time.Duration, max int) (*Limiter, error) {
	if max < 0 || totalTime <= 0 || tick <= 0 || tick > totalTime || totalTime%tick != 0 {
		return nil, errors.New("parameter error")
	}

	limiter := &Limiter{
		lock:      sync.Mutex{},
		totalTime: totalTime,
		tick:      tick,
		window:    make([]int64, int(totalTime/tick)),
		stop:      make(chan struct{}, 1),
		max:       int64(max / int(totalTime/tick)),
	}

	go limiter.Shift()

	return limiter, nil
}

//每隔一定时间移动一次窗口
func (l *Limiter) Shift() {
	ticker := time.NewTicker(l.tick)
I:
	for {
		select {
		case <-ticker.C:
			l.slidingWindow()
		case <-l.stop:
			break I
		}
	}
}

//移动窗口
func (l *Limiter) slidingWindow() {
	l.index++
	if l.index >= len(l.window) {
		l.index = 0
	}
	l.window[l.index] = 0
}

//添加计数
func (l *Limiter) Add() bool {
	defer l.lock.Unlock()
	l.lock.Lock()
	if l.window[l.index] >= l.max {
		return true
	}
	l.window[l.index]++
	return false
}

//停止计数
func (l *Limiter) Stop() {
	l.once.Do(func() {
		l.stop <- struct{}{}
	})
}
