package main

import "fmt"

func greeting(name string) string {
	return "Hello" + name
}
func add(value int) int {
	return 9 + value
}
func getmuliples(num1, num2 int) int {
	return num1 * num2
}
func main() {
	fmt.Println(greeting("Sravan"))
	fmt.Println(add(0))
	fmt.Println(getmuliples(3, 4))
}
