package main

import "fmt"

type Car struct {
	Model  string
	Year   int
	Engine Engine
}

type Engine struct {
	Serial string
	Type   string
}

func main() {

	var car1 Car
	car1.Model = "SUBARU"
	car1.Year = 1999
	car1.Engine.Serial = "1234567"
	fmt.Println(car1)

}