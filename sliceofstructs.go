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

	var car2 Car
	car2.Model = "TOYOTA"
	car2.Year = 2002
	car2.Engine.Serial = "432234"

	var car3 Car
	car3.Model = "MERCEDiS"
	car3.Year = 1995
	car3.Engine.Serial = "0101010"

	allcars := []Car{car1, car2}
	allcars = append(allcars, car3)

	fmt.Println(allcars)

}
