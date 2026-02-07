package main

import "fmt"

type Student struct {
	name, secondname string
	age, class       int
}

func main() {
	var std1 Student
	std1.name = "ANDREW"
	fmt.Println(std1.name)
	fmt.Println(std1.details())   // value receiver
	fmt.Println(std1.details2())  // pointer receiver
	//fmt.Println(std1.name, &std1.name)
}

// Value receiver
func (std Student) details() string {
	std.name= "KELVIN"
	return std.name
}

// Pointer receiver
func (stdp *Student) details2() string {
	stdp.name ="KEN"
	return stdp.name
}
