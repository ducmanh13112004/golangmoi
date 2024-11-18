package main

import "fmt"

// func (receiver ReceiverType) MethodName(parameters) returnType {
//     // Code
// }
// receiver: Là tham số đặc biệt đại diện cho đối tượng mà method được gọi.
// ReceiverType: Là kiểu dữ liệu (struct) mà method được gắn với.
// MethodName: Tên của method.
// parameters: Tham số truyền vào (nếu có).
// returnType: Kiểu dữ liệu trả về (nếu có).

type BankAccount struct {
	AccountNumber string
	Balance       float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.Balance += amount
		fmt.Printf("Nap tien thanh cong  ! so du moi cua ban la : %.2f\n", b.Balance)
	} else {
		fmt.Println("So tien nao phai lon hon 0 !")
	}
}

//rut tien
func (b *BankAccount) Withdraw(amount float64) {
	if amount > 0 {
		if amount <= b.Balance {
			b.Balance -= amount
			fmt.Printf("Rút tiền thành công! Số dư còn lại: %.2f\n", b.Balance)
		} else {
			fmt.Println("Số dư không đủ để rút.")
		}
	} else {
		fmt.Println("So tien phai lon hon  0 ! ")
	}
}

//Kiem tra so du
func (b BankAccount) CheckBalance() {
	fmt.Printf("Số dư hiện tại: %.2f\n", b.Balance)
}

func main() {
	account := BankAccount{AccountNumber: "123456789",
		Balance: 1000.0, // Số dư ban đầu
	}

	fmt.Printf("Tài khoản: %s, Số dư ban đầu: %.2f\n", account.AccountNumber, account.Balance)

	account.Deposit(500)   // Nạp thêm 500
	account.Withdraw(300)  // Rút 300
	account.Withdraw(1500) // Rút quá số dư
	account.CheckBalance() // Kiểm tra số dư
}
