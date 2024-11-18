package main

import (
	"fmt"
)

func main() {
	fmt.Println("Kết quả chương trình ")
	fmt.Println("---Goroutine---")
	Goroutine()
	fmt.Println("---channels---")
	channels()
	fmt.Println("---Select---")
	testselect()
	fmt.Println("---waitGroup---")
	testwaitGroup()
}
