package main

import "fmt"

// creating functions in go is pretty similar to creating functions in swift
// only differnce is that in swift you use colon (:)
//before giving the type of data the parameter will hold and what the return type will be
func add(a int, b int) int {
	return a + b
}

// we can use this syntax which is useful because it tells that all of these parameters
// have the same data type all of parameters are int
func add3(a, b, c int) int {
	return a + b + c
}

func main() {
	ans := add(1, 1)
	fmt.Println(ans)

	ans2 := add3(1, 1, 1)
	fmt.Println(ans2)
}
