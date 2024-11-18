//Tính tổng và giá trị trung bình của các phần tử trong một mảng
package main

import "fmt"

func testArray() {
	arr := [...]float64{1, 2, 3, 4, 5}
	fmt.Println("chiều dài của mảng: ", len(arr))
	//khai báo mảng không cần size
	arr2 := [...]string{"một", "hai", "ba"}
	fmt.Println("Kết quả : ", arr2)
	//mảng 2 chiều
	arr3 := [4][2]int{
		{1, 2},
		{1, 2},
		{1, 2},
		{1, 2},
	}
	fmt.Println("Kết quả mảng 2 chiều : ", arr3)
	fmt.Println("---------")
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arr3[i][j], "  ")
		}
		fmt.Println()
	}
}
