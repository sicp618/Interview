package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	RunContext2()
}

func RunContext2() {
	// 创建根 context
	rootCtx := context.Background()

	// 创建第一个 context，超时时间为 5 秒
	ctx1, cancel1 := context.WithTimeout(rootCtx, 5*time.Second)
	defer cancel1()

	var wg sync.WaitGroup
	wg.Add(1)
	go runC1(ctx1, context.Background(), &wg)

	wg.Wait()
}

func runC1(ctx context.Context, parentCtx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("runC1", time.Now().Format("2006-01-02 15:04:05"))
	select {
	case <-ctx.Done():
		fmt.Println("runC1 stop", time.Now().Format("2006-01-02 15:04:05"), ctx.Err())
	case <-time.After(7 * time.Second):
		fmt.Println("操作成功")
	}

	nextContext, _ := context.WithTimeout(parentCtx, 10*time.Second)
	wg.Add(1)
	go runC2(nextContext, wg)
}

func runC2(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("runC2", time.Now().Format("2006-01-02 15:04:05"))

	select {
	case <-ctx.Done():
		fmt.Println("runC2 stop", time.Now().Format("2006-01-02 15:04:05"), ctx.Err())
	case <-time.After(70 * time.Second):
		fmt.Println("操作成功")
	}
}

func RunContext1() {
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

func workerWaitDeath(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Goroutine %d 启动\n", id)

	select {
	case <-ctx.Done():
		fmt.Printf("Goroutine %d 收到取消信号: %v\n", id, ctx.Err())
	}
}

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Goroutine %d 启动\n", id)

	select {
	case <-ctx.Done():
		fmt.Printf("Goroutine %d 收到取消信号: %v\n", id, ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Printf("Goroutine %d 完成任务\n", id)
	}
}
