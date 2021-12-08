package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

var s = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	// TryAcquire ... 他のgoroutineを待たずに終了
	if isAcquire := s.TryAcquire(1); !isAcquire {
		fmt.Println("Could not get lock.")
		return
	}
	// Acquire ... 他のgoroutineを待ってから実行
	//if err := s.Acquire(ctx, 1); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(2 * time.Second)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}
