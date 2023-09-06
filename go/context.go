package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// 创建一个父 context
	parentCtx := context.Background()

	// 创建一个带有取消操作的 context，10 秒后取消
	ctx, cancel := context.WithTimeout(parentCtx, 10*time.Second)
	defer cancel()

	// 使用 WaitGroup 来等待所有 Goroutine 完成
	var wg sync.WaitGroup

	// 启动 3 个 Goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	// 等待所有 Goroutine 完成
	wg.Wait()

	fmt.Println("所有 Goroutine 已完成")
}

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Printf("Goroutine %d 收到取消信号: %v\n", id, ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Printf("Goroutine %d 完成任务\n", id)
	}
}
