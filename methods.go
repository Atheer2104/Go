package main

import "fmt"

type rect struct {
	width, height int 
}

// creating a method on the struct rect named area and it will return and int
func (r rect) area() int {
	return r.height * r.width
}

// pointer of rect struct the above is a value type reciver
func (r *rect) perim() int {
	return 2 * r.width * 2 * r.height
}

func main() {

	r := rect {width: 10, height: 5}
	fmt.Println("Area:", r.area())

	r_ptr := &r
	fmt.Println("perim:", r_ptr.perim())

}