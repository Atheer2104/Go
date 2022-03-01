package main

import "fmt"

// creating a function that will return anonymous function
func return_anoyn_func() func(string) {
	// returning an anonymous function 
	return func(msg string){
		fmt.Println(msg)
	}
}

// returning a func and that func which we will return will return an int 
func int_seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// anonymous function, a function with out a name
	func (msg string) {
		fmt.Println(msg)
	}("Hello Anonymous")

	// we are storing the anon func to the varible
	print_func := return_anoyn_func()
	print_func("Hello returned from anonymous")

	next_int := int_seq()

	// i gets increamented by one, it's like it's mean remembered and that is because of closures
	fmt.Println(next_int())
	fmt.Println(next_int())

}