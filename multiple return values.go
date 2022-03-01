package main 

import "fmt"

// here we are telling it we will return two thing and both of them are type of int 
func add_sub(a, b int)(int, int) {
	// here we are returning the two diffierent kind of values and seperating them by a coma (,)
	return a+b, a-b
}

func main() {

	add_res, sub_res := add_sub(1, 1)
	fmt.Println("add_res:", add_res, "sub_res:", sub_res)

}