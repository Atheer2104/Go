package main 

import "fmt"

// here we are creating our own variadic functions 
//we have a parameter name followed by three dots ...
// and the data type of our int 
func mult(nums...int)int {
	total := 1
	for _, num := range nums {
		 total *= num
	}
	return total
}

func main() {
	// Println is an example of a variadic function.
	// there is no limit to how many we parameters we can have
	fmt.Println("this", "is", "an", "example")

	// we can call mult with as many parameters as we like
	fmt.Println(mult(1, 2, 3, 4))

	// variadic functions can also be applied to slices:
	nums := []int {1, 2, 3, 4}
	fmt.Println(mult(nums...))
}