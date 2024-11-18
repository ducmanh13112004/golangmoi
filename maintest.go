package main

import "fmt"

func main() {
	// Khởi tạo slice với các phần tử ban đầu
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice ban đầu:", numbers)
	fmt.Println("Độ dài:", len(numbers))
	fmt.Println("Dung lượng:", cap(numbers))

	// Thêm phần tử vào slice bằng cách sử dụng append
	numbers = append(numbers, 6)
	fmt.Println("Slice sau khi thêm phần tử:", numbers)

	// Truy cập phần tử của slice
	fmt.Println("Phần tử đầu tiên:", numbers[0])
	fmt.Println("Phần tử thứ hai:", numbers[1])

	// Cắt một phần của slice (slicing)
	subSlice := numbers[1:4] // Lấy các phần tử từ vị trí 1 đến 3 (không bao gồm vị trí 4)
	fmt.Println("Slice con:", subSlice)
	fmt.Println("hi")
}
