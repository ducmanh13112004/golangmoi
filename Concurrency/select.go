package main

import "fmt"

func testselect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c1 <- "From c1"
	}()

	go func() {
		c2 <- "From c2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
