package main

import "fmt"

func main() {

	var counter int = 5

	var message string // default value ""
	message = "Hello"

	var factor float32 // default value 0
	factor = 4.2

	var enabled bool // false

	fmt.Println(counter)
	fmt.Println(message)
	fmt.Println(factor)
	fmt.Println(enabled)

	fmt.Println("7/2 =", 7/2)
	fmt.Println("7/2.0 =", 7/2.0)
	fmt.Println("7.2/2 =", 7.2/2)
}
