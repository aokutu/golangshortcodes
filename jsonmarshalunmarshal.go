package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	var person1 Person
	person1.Name = "ANDREW"
	person1.Age = 35

	fmt.Println(person1)

	x, _ := json.Marshal(person1)

	fmt.Println(string(x))

	var Person2 Person

	err := json.Unmarshal([]byte(string(x)), &Person2)
	if err != nil {
		fmt.Println("GOT AN  ERROR")
	}

	fmt.Println(Person2)

}
