package main

import "fmt"

func sum(arr []int, c chan int) {
	total := 0
	for _, v := range arr {
		total += v
	}
	c <- total // Gửi kết quả vào channel
}
func testconcurency() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	mid := len(numbers) / 2
	part1 := numbers[:mid] // Phần đầu tiên
	part2 := numbers[mid:] // Phần còn lại
	//tao chan
	c := make(chan int)
	//tinh tong song song
	go sum(part1, c)
	go sum(part2, c)
	// nhận kết quả từ 2 goroutines
	sum1 := <-c
	sum2 := <-c

	fmt.Printf("Tổng : %d\n", sum1+sum2)
}
