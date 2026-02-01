package main

import "fmt"

type Student struct {
	name, secondname string
	age, class       int
}

func main() {

	var std1 Student
	std1.name = "ANDREW"
	fmt.Println(std1.details())

}

//a method  is a function  with  a type  preceeding the  function name as shown  below 
func (std Student) details() string {

	return std.name
}