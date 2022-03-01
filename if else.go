package main

import "fmt"

func main() {

	i := 5

	// if and else statment in python
	if i%2 == 0 {
		fmt.Println(i, "i is even.")
	} else {
		fmt.Println(i, "is odd.")
	}

	k := 100
	if k == 100 {
		fmt.Println(k, "is 100")
	}

	// a else if statment in go very similar to other languages
	j := 25
	if j < 50 {
		fmt.Println(j, "is less than 50")
	} else if j > 50 {
		fmt.Println(j, "is greater than 50")
	} else {
		fmt.Println(j, "is equal to 50")
	}

	// there is no tern operator in go or nil colesing in go
	// a ? b : c

}
