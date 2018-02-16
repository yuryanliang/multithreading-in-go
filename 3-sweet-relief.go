package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2

	operationDone := make(chan bool)
	go func() {
		b = a * b

		operationDone <- true
	}()

	<-operationDone

	a = b * b

	fmt.Println("Hit Enter when you want to see the answer")
	fmt.Scanln()

	fmt.Printf("a = %d, b = %d\n", a, b)
}
