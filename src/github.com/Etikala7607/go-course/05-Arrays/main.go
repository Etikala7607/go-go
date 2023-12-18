package main

import "fmt"

func main() {
	// Arrays
	//var fruitArra [2]string

	// //Assign values
	// fruitArra[0] = "Apple"
	// fruitArra[1] = "Mango"

	//declare and assign
	// fruitArra := [2]string{"mango", "orange"}

	// fmt.Println(fruitArra)
	// fmt.Println(fruitArra[1])

	fruitSlice := []string{"mango", "apple", "grapes", "banana"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
	for i := 0; i < len(fruitSlice)-1; i++ {
		fmt.Println(i)
	}
	for i := len(fruitSlice) - 1; i >= 0; i-- {
		fmt.Println("the reverse order is " + fmt.Sprintf("%d", i))
	}
}
