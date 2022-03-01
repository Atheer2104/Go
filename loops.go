package main

// go has only one type loop and it's for while other langauges has while loop for example

import "fmt"

func main() {

	i := 0
	// creating a for loop in go
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	// also a for loop but this is the common we are used to see in other languages like javascript
	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}

	for k := 0; k <= 10; k++ {
		if k%2 == 0 {
			// does the same as the other languages just skips the code that would otherwise happen
			continue
		}

		fmt.Println(k)
	}

	// however if we want some like while we could do for { } and we then use break to cancel it else would
	// keep running for ever

}
