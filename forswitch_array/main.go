package main

import "time"
import "fmt"

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	// fmt.Print("Kết quả chương trình")
	testArray()
	testSwitchCase()
	ifElse()
}
