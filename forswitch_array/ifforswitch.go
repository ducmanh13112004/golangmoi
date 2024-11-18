package main

import "fmt"
func ifElse(){
	 Age :=18
	if Age>=18{
		fmt.Println("Bạn đã đủ 18 tuổi !")
	}else{
		fmt.Println("Bạn chưa đủ 18 tuổi !")
	}
	//else không đươc cách hàng xuống (phải viết liền kề sau dấu  } của If)
}
func testSwitchCase() {
	var weight, height float64
	var chocie string
	for {
		fmt.Print("Nhập cân nặng (Kg): ")
		fmt.Scan(&weight)

		fmt.Print(("Nhập chiều cao (m): "))
		fmt.Scan(&height)

		bmi := weight / (height * height)

		fmt.Printf("Chỉ số BMI của bạn là:  %.2f\n", bmi)

		switch {
		case bmi < 18.5:
			fmt.Println("Thiếu cân !")
		case bmi >= 18.5 && bmi < 24.9:
			fmt.Println("Phân loại: Bình thường")
		case bmi >= 25 && bmi < 29.9:
			fmt.Println("Phân loại: Thừa cân")
		default:
			fmt.Println("Phân loại: Béo phì")
		}

		fmt.Print(" Bạn có muốn tính lại cho người khác không: ")
		fmt.Scan(&chocie)

		if chocie != "y" {
			fmt.Println("Chương trình kết thúc .")
			break
		}
	}
	
}



// package main
// import "fmt"

// func main(){
// 	fmt.Println("Hello Word")
// 	sum :=0
	
// 	for i := 1; i <= 100; i++ {
// 		sum += i
// 		// fmt.Println("Tổng các số từ 1 đến 100 là:", sum) //bỏ vào vòng for thì in ra 100 lần 
// 	}
// 	fmt.Println("Tổng các số từ 1 đến 100 là:", sum)
// }