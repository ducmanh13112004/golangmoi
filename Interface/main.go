package main

import "fmt"

// Định nghĩa interface Vehicle
type Vehicle interface {
	Move() string     // Mô tả cách phương tiện di chuyển
	MaxSpeed() int    // Tốc độ tối đa
	FuelType() string // Loại nhiên liệu
}

// Struct Car (Ô tô)
type Car struct {
	Brand string
}

func (c Car) Move() string {
	return "Ô tô di chuyển trên đường bằng 4 bánh."
}

func (c Car) MaxSpeed() int {
	return 200 // km/h
}

func (c Car) FuelType() string {
	return "Xăng"
}

// Struct Motorcycle (Xe máy)
type Motorcycle struct {
	Brand string
}

func (m Motorcycle) Move() string {
	return "Xe máy di chuyển trên đường bằng 2 bánh."
}

func (m Motorcycle) MaxSpeed() int {
	return 150 // km/h
}

func (m Motorcycle) FuelType() string {
	return "Xăng"
}

// Struct Airplane (Máy bay)
type Airplane struct {
	Model string
}

func (a Airplane) Move() string {
	return "Máy bay bay trên không trung."
}

func (a Airplane) MaxSpeed() int {
	return 900 // km/h
}

func (a Airplane) FuelType() string {
	return "Nhiên liệu máy bay"
}

// Hàm nhận một Vehicle và hiển thị thông tin
func DescribeVehicle(v Vehicle) {
	fmt.Println(v.Move())
	fmt.Printf("Tốc độ tối đa: %d km/h\n", v.MaxSpeed())
	fmt.Printf("Loại nhiên liệu: %s\n", v.FuelType())
	fmt.Println("-----------------------")
}
func main() {
	// Tạo các phương tiện
	car := Car{Brand: "Toyota"}
	motorcycle := Motorcycle{Brand: "Honda"}
	airplane := Airplane{Model: "Boeing 747"}

	// Tạo danh sách phương tiện
	vehicles := []Vehicle{car, motorcycle, airplane}

	// Hiển thị thông tin từng phương tiện
	for _, vehicle := range vehicles {
		DescribeVehicle(vehicle)
	}
}
