package main

import "fmt"

func main() {
	x := 15
	y := 10

	// if else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", x, y)
	}

	// else if

	color := "white"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("color is neither red nor blue")
	}

	//switch

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("Color is blue")

	case "white":
		fmt.Println("color is neither blue nor white")
	}
}
