package main 

import "fmt"

func main() {
	val := 20
	fmt.Println(val)

	// ptr is set to the adress of where val is stored (refercing)
	// using this symbol & to get the adress 
	ptr := &val

	// printing the adress of where value 20 is stored 
	fmt.Println(ptr)

	//print the value stored at the adress aka (dereference)
	// in this we are using the * symbol to dereference
	fmt.Println(*ptr)

	// changing the value of ptr and not the adress but ptr and val 
	// have the same adress and when I change the value on ptr it 
	// also changes the value on val because the have the same adress and that 
	// adress has a value 
	*ptr = 10
	fmt.Println(*ptr)
	fmt.Println(val)

	val = 50 
	fmt.Println(*ptr)
}