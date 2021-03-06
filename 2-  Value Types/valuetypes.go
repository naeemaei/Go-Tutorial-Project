package main

import "fmt"

const uploadFolder string = "Iran" // Declare global const in golang

func main() {

	// Declare local const in golang, when we have two global and local const with same name, global const was ignored.
	const uploadFolder string = "Persepolis"
	const num = 10 - 1

	//Sample integer
	var counter int = 5

	// Default value ""
	var message string
	message = "Hello"

	// Default value 0
	var factor float32
	factor = 4.2

	//Other way to declare variable
	appName := "Hello Iran"

	// Default value false
	var enabled bool

	// Define two variables
	var firstName, lastName string = "hamed", "naeemaei"

	// Multiple Declarations
	var (
		id                    int    = 1
		cityName, countryname string = "Chalus", "Iran"
	)

	// print variable
	fmt.Println(counter)
	fmt.Println(message)
	fmt.Println(factor)
	fmt.Println(enabled)
	fmt.Println(appName)
	fmt.Println(firstName, lastName)
	fmt.Println(id, cityName, countryname)

	// calculation in print
	fmt.Println("7/2 =", 7/2)
	fmt.Println("7/2.0 =", 7/2.0)
	fmt.Println("7.2/2 =", 7.2/2)

	//print constants
	fmt.Println(uploadFolder)
	fmt.Println(num)
}
