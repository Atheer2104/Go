package main

import "fmt"

func main() {

	// maps in go similar to dictionary

	// specify key/value pair in go
	// in this case the key is string and value is int
	m := make(map[string]int)

	m["a"] = 0
	m["b"] = 1

	fmt.Println(m)

	// printing value of a specific key
	fmt.Println(m["a"])

	// len() function is overloaded to work with maps:
	fmt.Println(len(m))

	// remove key/pair from map. (which can be done with delete keyword)
	delete(m, "a")
	fmt.Println(m)

	// initalizing map with another way when we know the key/pair values
	m2 := map[string]int{"key1": 1, "key2": 2}
	fmt.Println(m2)

	// the value and whether the value is present
	// will return the value and val will be assignd to that and the second value will be
	// whether it's present in the dictionary
	val, is_val_present := m["b"]
	fmt.Println(val)
	fmt.Println(is_val_present)

	// when we don't care of value then we use _
	_, is_val_present2 := m["a"]
	fmt.Println(is_val_present2)
}
