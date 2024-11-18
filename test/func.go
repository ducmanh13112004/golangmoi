package main

import "fmt"

// Hàm tính tổng
func add(a int, b int) int {
	return a + b
}

func rectangle(width, height float64) (float64, float64) {
	area := width * height
	perimeter := 2 * (width + height)
	return area, perimeter
}

func functest() {
	fmt.Println("Tổng của 3 và 4 là:", add(3, 4))

	area, perimeter := rectangle(5.0, 3.0)
	fmt.Println("Diện tích:", area)
	fmt.Println("Chu vi:", perimeter)
}

func main() {
	functest()
}
