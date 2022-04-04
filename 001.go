package main

import "fmt"

func main() {
	fmt.Println("Hi my name is vivek")
	foo()
	for i := 0; i <= 5; i++ {
		fmt.Println("count ", i)
	}
}

func foo() {
	fmt.Println("Yo yo yo")
}
