package main

import "fmt"

// Khai báo biến bên ngoài hàm
var a int = 10         // Khai báo biến kiểu int
var b float64 = 20.5   // Khai báo biến kiểu float64
var c string = "Hello" // Khai báo biến kiểu string
var d bool = true      // Khai báo biến kiểu boolean
var e complex64 =10+1i
var e1 complex64 =10+1i
var r rune = 'あ' // đại diện cho một ký tự Unicode (có thể làm việc với các ký tự phức tạp hoặc đa ngôn ngữ )
func varibales() {
	// Khai báo biến bên trong hàm (thì mới sử dụng cú pháp ngắn gọn được :=)
	x := 30
	y := 20.6
	z := "World"
	typeData :=4
	typeData2 :=5.6
	isStudent := false

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("KQ:", e + e1)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("Giá trị rune:", r)             // In mã Unicode của ký tự
    fmt.Printf("Ký tự tử rune: %c\n", r) // Chuyển đổi mã rune thành ký tự
	fmt.Println("kết quả phép toán: ",float64(typeData)+typeData2) //ép qua kiểu float để thực hiện phép toán
	fmt.Println("isStudent:", isStudent)
}
