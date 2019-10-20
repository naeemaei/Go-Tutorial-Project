/// Lesson 1 : Hello World!

package main // The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library

// import necessary libraries
import (
	"fmt"
)

// Golang function
func hello() {

	fmt.Println("hello World!")
}

// main function: start of project
func main() {

	hello()
}

// How run codes
