package main

import (
	"fmt"
	"sync"
	"time"
)

type SlidingWindow struct {
	Buckets  map[int64]*counterBucket
	Mu       *sync.RWMutex
	Duration time.Duration
}

type counterBucket struct {
	count int64
}

func NewSlidingWindow(duration time.Duration) *SlidingWindow {
	return &SlidingWindow{
		Buckets:  make(map[int64]*counterBucket),
		Mu:       &sync.RWMutex{},
		Duration: duration,
	}
}

func (sw *SlidingWindow) getCurrentBucket() *counterBucket {
	now := time.Now().Unix()
	var bucket *counterBucket
	var ok bool
	if bucket, ok = sw.Buckets[now]; !ok {
		bucket = &counterBucket{}
		sw.Buckets[now] = bucket
	}
	return bucket
}

func (sw *SlidingWindow) removeExpiredBucket() {
	now := time.Now().Unix() - int64(sw.Duration.Seconds())
	for timestamp := range sw.Buckets {
		if timestamp <= now {
			delete(sw.Buckets, timestamp)
		}
	}
}

func (sw *SlidingWindow) Sum(now time.Time) int64 {
	sum := int64(0)

	sw.Mu.Lock()
	defer sw.Mu.Unlock()

	for timestamp, bucket := range sw.Buckets {
		if timestamp >= time.Now().Unix()-int64(sw.Duration.Seconds()) {
			sum += bucket.count
		}
	}

	return sum
}

func (sw *SlidingWindow) Avg(now time.Time) int64 {
	return sw.Sum(now) / int64(sw.Duration.Seconds())
}

func (sw *SlidingWindow) Increment() {
	sw.Mu.Lock()
	defer sw.Mu.Unlock()

	b := sw.getCurrentBucket()
	b.count++
	sw.removeExpiredBucket()
}

func main() {
	sw := NewSlidingWindow(10 * time.Second)
	for range make([]int, 10000) {
		sw.Increment()
		time.Sleep(1 * time.Millisecond)
	}

	fmt.Println(sw.Avg(time.Now()))
}
