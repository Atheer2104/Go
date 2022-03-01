package main

import "fmt"

func main() {

	// creating an int array that will hold in this case 5 elements
	// when creating an empty int array like this it will hold a 5 elements
	// and all of them will be zeros (0) this is how it works in go
	var int_arr [5]int
	fmt.Println(int_arr)

	// in an empty bool array all the values will be false
	var bool_arr [10]bool
	fmt.Println(bool_arr)

	// modifying the array
	int_arr[0] = 5
	int_arr[1] = 10

	// creating an array and directly assigning values to it
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	// length of array
	fmt.Println(len(a))

	// creating a two dimensional array
	var aa [5][5]int
	fmt.Println(aa)

	count := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			aa[i][j] = count
			count++
		}
	}
	fmt.Println(aa)

}
