package main 

import "fmt"

// a recursive function is a func that 
// calls it self until a base case is met
func fact(n int)int {
	// base case
	if n <= 1 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(26))
}