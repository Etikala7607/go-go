package main

import (
	"fmt"
)

func main() {
	a := 5
	var ptr *int
	ptr = &a

	fmt.Println("value is ", *ptr)
	*ptr = 20

	fmt.Println("modified value is ", a)
	fmt.Printf("%T\n", a)
	// use * to read the value
	fmt.Println(*&a)

}
