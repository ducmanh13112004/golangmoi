package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Đánh dấu goroutine hoàn thành
	fmt.Printf("Worker %d bắt đầu\n", id)
	fmt.Printf("Worker %d hoàn thành\n", id)
}

func testwaitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Thêm 1 goroutine
		go worker(i, &wg)
	}

	wg.Wait() // Chờ tất cả goroutines hoàn thành
	fmt.Println("Tất cả workers đã hoàn thành.")
}
