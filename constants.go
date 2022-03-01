package main

import "fmt"

func main() {

	// any object can be a constant ie string, int, bool and float32
	// a constant can not be changed once it value is set then can not be changed
	// go doesn't have let so const is the go version of let
	const pi float64 = 3.1415926

	fmt.Println(pi)

	const c int = 1000

	fmt.Println(c)

}
