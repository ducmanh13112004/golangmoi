// Quản lý điểm số học sinh
package main

import (
	"fmt"
	"sort"
)

func QLDSSV() {
	var n int
	fmt.Println("Nhập số lượng sinh viên: ")
	fmt.Scan(&n)
	//khởi tạo mảng
	scores := make([]int, n)
	//nhập điểm số
	for i := 0; i < n; i++ {
		fmt.Printf("Nhập điểm của sinh viên thứ  %d: ", i+1)
		fmt.Scan(&scores[i])
	}
	//tìm max,min,tổng và số học sinh đậu
	max, min, sum, pass := scores[0], scores[0], 0, 0
	for _, score := range scores {
		if score > max {
			max = score
		}
		if score < min {
			min = score
		}
		sum += score
		if score >= 50 {
			pass++
		}
	}
	average := float64(sum) / float64(n)
	//sap xep mang giam dan
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))

	//in ra ket qua
	fmt.Printf("Diem cao nhat la: %d\n ", max)
	fmt.Printf("Điểm thấp nhất: %d\n", min)
	fmt.Printf("Diem trung binh : %.2f\n", average)
	fmt.Printf("Số học sinh đậu: %d\n", pass)
	fmt.Printf("Mảng điểm sau khi sắp xếp: %v\n", scores)

	//tim kiếm điểm
	var target int
	fmt.Println("Nhap diem can tim kiem : ")
	fmt.Scan(&target)
	found := false
	for i, score := range scores {
		if score == target {
			fmt.Printf("Diem %d nam o vi tri: %d\n", target, i)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Khong tim thay diem nay !")
	}
}
