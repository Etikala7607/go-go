package main

import "fmt"

func main() {
	ids := []int{1, 3, 5, 89, 2, 76}
	//loop through ids
	for i, id := range ids {
		fmt.Printf("%d-ID: %d*2\n", i, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d*2\n", id)
	}
	//Add all ids
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// Range with map
	emails := map[string]string{"Ravan": "ravan@gmaail.com", "Anamoly": "anamoly@gmail.com", "pinky": "pinky@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	for k := range emails {
		fmt.Println("Name:" + k)
	}
	length := len(emails)
	fmt.Println(length)
	keys := make([]string, len(emails))
	i := 0
	for k := range emails {
		keys[i] = k
		i++
	}
	for j := len(keys) - 1; j >= 0; j-- {
		k := keys[j]
		v := emails[k]
		fmt.Printf("%s: %s\n", k, v)
	}
}
