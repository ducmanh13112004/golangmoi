package main

import "fmt"

func workerchan(c chan string) {
	c <- "Hello from worker" // Gửi dữ liệu vào channel
}

func channels() {
	ch := make(chan string) // Tạo channel

	go workerchan(ch) // Chạy worker trong goroutine
	msg := <-ch       // Nhận dữ liệu từ channel
	fmt.Println(msg)  // In ra "Hello from worker"
}
