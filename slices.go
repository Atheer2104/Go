package main

import "fmt"

func main() {
	// creating a slice of type int array and the size of slice is 3
	// slices are similar to an array
	s := make([]int, 3)

	s[0] = 1
	s[1] = 2
	s[2] = 3

	// the similarites
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println(len(s))

	// append function is unique to slices
	// here we are telling we are appending an int into s and what that value will be
	s = append(s, 4)

	fmt.Println(s)

	// we can also append multiple values at the same time
	s = append(s, 5, 6, 7)

	fmt.Println(s)

	// slice syntax

	// we are getting the values starting from and finishing at 3 however we will not print
	// the 3 value onlu 1 and 2
	fmt.Println(s[1:3])

	// here we are telling it to start from the first value ie 0 and ending when it reach 3
	fmt.Println(s[:3])

	// here we are starting at the given index and going all the way through until the end of slice
	fmt.Println(s[1:])

	// concise slice defination

	//creating a slice with values
	t := []int{100, 200, 300}
	fmt.Println(t)

	// x := s
	// x[0] = 500
	// earlier we said that x was the same value as s as when we modified x
	// it modified s too and that is because of pointers, good thing to know
	// they are entertwined/entangled
	// fmt.Println(x)
	// fmt.Println(s)

	// if we want to prevent from changing both x and s then we use copy
	x := make([]int, len(s))
	// here we are copying s and pasting it to x
	copy(x, s)

	x[0] = 500
	fmt.Println(x)
	fmt.Println(s)

	// 2-D slice, similar to arrays though lengths of slice may vary.
	ss := make([][]int, 3)
	// [ [0], [1,2], [2, 3, 4] ]
	for i := 0; i < 3; i++ {
		inner_len := i + 1
		ss[i] = make([]int, inner_len)
		for j := 0; j < inner_len; j++ {
			ss[i][j] = i + j
		}
	}
	fmt.Println(ss)

}
