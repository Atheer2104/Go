package main

import "fmt"

// creating a struct in go
type employee struct {
	first_name string
	last_name  string
	id         int
}

func main() {

	// two ways of creating an employee struct if we want to see the parameter names or not
	// and if we forget to give a value to the struct properties
	// then we will get and empty version of the data type "" (string) 0 (int)
	fmt.Println(employee{first_name: "Homer", last_name: "Simpson", id: 1})
	fmt.Println(employee{"Whalen", "Smithers", 2})

	fmt.Println(employee{first_name: "Frank", last_name: "Grimes"})

	emp := employee{first_name: "Montgomery", last_name: "Burns", id: 4}

	//  access the properties of an employee object
	fmt.Println(emp.first_name)
	fmt.Println(emp.last_name)
	fmt.Println(emp.id)

	emp_ptr := &emp
	fmt.Println(emp_ptr.first_name)

	// pointors also works with structs

	emp_ptr.first_name = "Marge"
	fmt.Println(emp_ptr.first_name)
	fmt.Println(emp.first_name)

}
