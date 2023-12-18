package main

import "fmt"

func main() {
	fmt.Println("hello Go!")

	// long method
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// //short method
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Number is %d\n", i)
	// }
	//fizbuzz
	for me := 1; me <= 100; me++ {
		if me%15 == 0 {
			fmt.Println("Fizbuzz")
		} else if me%3 == 0 {
			fmt.Println("fizz")
		} else if me%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(me)
		}
	}

}
