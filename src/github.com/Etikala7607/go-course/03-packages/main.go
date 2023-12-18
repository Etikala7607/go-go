package main

import (
	"fmt"
	"math"

	"example.com/mymodule/03-packages/strutil"
)

func main() {
	fmt.Println(math.Floor(8.8))
	fmt.Println(math.Ceil(8.6))
	fmt.Println(strutil.Reverse("navaR"))

}
