package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

type Triangle struct {
	height int
	width  int 

}

func main() {

	var r1 rectangle
	r1.length = 5
	r1.width = 4
	var r2 Triangle
	r2.width = 4
	r2.height = 5
	fmt.Println("RECT",Areax(r1))

	fmt.Println("TRI",Trireax(r2))

}


func Areax(rec rectangle) int {
	return rec.width * rec.length
}

func Trireax(tri Triangle) int {
	return (tri.width * tri.height) / 2
}
