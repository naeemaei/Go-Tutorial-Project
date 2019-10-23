package main

import "fmt"

func main() {
	// 1: Simple version of for loop works as c,c++,c#,java and etc .
	sum := 0
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Print(i * j)
			fmt.Print(" ")
		}
		fmt.Println(" ")

		sum += i
	}
	fmt.Println(sum) // 55 (1+2+3+4+5+6+7+8+9+10)

	// 2: While loop
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n) // 8 (1*2*2*2)

	// 3: Infinite loop

	/*
		sum = 0
		for {
			sum++ // repeated forever
		}
		fmt.Println(sum) // never reached
	*/

	// 4: For-each loop
	customersList := []string{"ali", "philip", "john", "fred", "shoji"}
	for i, s := range customersList {
		fmt.Println(i, s)
	}

	///////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////

	// Exit from loop
	sum = 0
	for i := 1; i < 10; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		if i > 8 {
			break

		}
		sum += i
	}
	fmt.Println(sum) // 20 (2+4+6+8)

}
