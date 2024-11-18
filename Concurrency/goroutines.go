package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(100 * time.Millisecond) // Tạm dừng 500ms
	}
}

func Goroutine() {
	go printMessage("Goroutine lần  1") // Chạy trong goroutine
	go printMessage("Goroutine lần  2") // Chạy trong goroutine

	// Goroutine chính
	printMessage("Main Goroutine ")
}
