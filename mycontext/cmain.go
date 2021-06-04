package mycontext

import (
	"context"
	"math/rand"
	"time"

	"fmt"
)

func ContextMain() {
	WithCancel()
	//WithTimeout()
	//WithDeadline()
	//WithValue()
}

// 从随机生成的数字中找到20，找到后通知主协程。
func subRoutine(id int, ctx context.Context, data chan int) {
	fmt.Println("entry")
	number := 0

	for {
		select {
		case <-ctx.Done():
			fmt.Println(id, "Exit")
			data <- -1
			return
		case <-time.Tick(1 * time.Second):
			number = rand.Intn(100)
			fmt.Println(id, "number", number)
			if number == 20 {
				fmt.Println(id, "Found....")
				data <- number
				return
			}
		}
	}
}

func WithCancel() {
	data := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		go subRoutine(i, ctx, data)
	}

	count := 0
	for d := range data {
		count += 1
		fmt.Println("Found", d)
		if count == 5 {
			break
		}
		if d != -1 {
			cancel()
		}
	}

	fmt.Println("WithCancel exit")
}

func WithTimeout() {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data := make(chan int)
	for i := 0; i < 5; i++ {
		go subRoutine(i, ctx, data)
	}

	count := 0
	for d := range data {
		count += 1
		fmt.Println("Found", d)
		if count == 5 {
			break
		}
	}

	fmt.Println("WithTimeout exit, elapsed", time.Since(start))
}

func WithDeadline() {
	deadline := time.Now().Add(5 * time.Second)
	start := time.Now()
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	data := make(chan int)
	for i := 0; i < 5; i++ {
		go subRoutine(i, ctx, data)
	}

	count := 0
	for d := range data {
		count += 1
		fmt.Println("Found", d)
		if count == 5 {
			break
		}
	}

	fmt.Println("WithDeadline exit, elapsed", time.Since(start))
}

func WithValue() {
	ctx := context.WithValue(context.Background(), "id", 11111111)
	data := make(chan int)
	go func(c context.Context, data chan int) {
		if id := c.Value("id"); id != nil {
			data <- id.(int)
			return
		}
	}(ctx, data)

	fmt.Println("ID", <-data)
}
