package main

import "fmt"

func main() {
	num int := {34,12,9,43,19}

	bubbleSort(num)
	print(num)
}

func bubbleSort(array []int) {
	for x := len(array) - 1; x > 0; x-- {
		for j := 0; j < x; j++ {
			if array[j] > array[j +1] {
				swap(array, j)
			}
		}
	}
}

func swap(array []int, index int) {
	temp := array[index]
	array[index] = array[index + 1]
	array[index + 1] = temp

}
