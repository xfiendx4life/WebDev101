package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type contexter interface {
	DoSmthLong(context.Context)
}

func DoSmthLong(ctx context.Context) {
	time.Sleep(2 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("contxt done")
		return
	default:
		fmt.Println("Job is Done")
	}

}

func check(ctxter contexter) {
	ctxter.DoSmthLong(context.Background())
}

func third() {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	var s int64
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// go func(i int) {
		// 	atomic.AddInt64(&s, int64(i))
		// 	wg.Done()
		// }(i)
		go func(i int) {
			mtx.Lock()
			defer mtx.Unlock()
			s += int64(i)
			wg.Done()
		}(i)
	}
	fmt.Println(s)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	DoSmthLong(ctx)
	_ = cancel
	ctx, cancel = context.WithCancel(context.Background())

	wg.Add(1)
	go func(ctx context.Context) {
		defer func() {
			fmt.Println("func 1 is done")
			wg.Done()
		}()
		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
	wg.Add(1)
	go func(ctx context.Context) {
		defer func() {
			fmt.Println("func 2 is done")
			wg.Done()
		}()
		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	wg.Add(1)
	go func(context.CancelFunc) {
		time.Sleep(3 * time.Second)
		cancel()
		wg.Done()
	}(cancel)

	wg.Wait()

}
