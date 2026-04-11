package main

import "fmt"

func main() {
	x := 5
	pointx := &x

	*pointx = 6
	fmt.Println(x)

}