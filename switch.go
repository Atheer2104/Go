package main

import "fmt"

func main() {
	
	i := 2

	// switch statment in go, a thing to notice with go switch statment is that the default 
	// statment is not required by us to implement if none of the cases in switch is called
	// then nothing will be fired
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")

	}


}