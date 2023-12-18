package main

import "fmt"

func main() {
	// define MAp
	// emails := make(map[string]string)

	// //Asssign KV

	// emails["Ravan"] = "ravan@gmail.com"
	// emails["Anamoly"] = "anamoly@gmail.com"
	// emails["pinky"] = "pinky@gmail.com"
	//decalre map and KV
	emails := map[string]string{"Ravan": "ravan@gmaail.com", "Anamoly": "anamoly@gmail.com", "pinky": "pinky@gmail.com"}
	fmt.Println(emails)
	fmt.Println(len(emails))

	//Delete  from map
	delete(emails, "pinky")
	fmt.Println(emails)

	// check if exists
	key, exists := emails["SRavan"]
	if exists {
		fmt.Println("email of ravan is %s\n", key)
	} else {
		fmt.Println("email not found in map")
	}
}
