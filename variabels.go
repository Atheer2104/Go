package main

import "fmt"

// instead of typing var all the time for creating many variables we can instead use this syntax below
/*var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)*/

// we can declare a varible with same name both instad the main func and outside (also called package level)
// also every variable has to be used you can't create a variable and don't use it
//var i int = 27

func main() {
	// when declaring variable we type name of variable and then the type and giving it a int in the case with
	// the equal operator
	i := 42
	// we could write the code on top a bit shorter with
	//j := 42
	// we could also write it like this
	// var k int
	//k = 42

	fmt.Println(i)
}
