package main

import "fmt"

func main() {

	// range is similar to range in python
	str_arr := []string{"a", "b", "c", "d"}

	// c is for character where we are looping over in the str array
	// i is just the index
	for i, c := range str_arr {
		fmt.Println("index", i)
		fmt.Println("c", c)
	}

	// can also range over key/values pairs ie maps aka dictionarys
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println("key:", k, "val:", v)
	}

	// maybe we only want the key or the index and don't care of the value
	// we could do this too
	for k := range m {
		fmt.Println("key:", k)
	}

}
