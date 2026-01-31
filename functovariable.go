package main

import "fmt"

func main() {

	var calculate func(float64, float64) float64
	calculate = calcarea
	output := calculate(5, 5)
	fmt.Println(output)
	calculate = calcperimeter
	output = calculate(5, 5)
	fmt.Println(output)

}

func calcarea(length, width float64) float64 {

	return length * width
}

func calcperimeter(length, width float64) float64 {

	return length + length + width + width
}
