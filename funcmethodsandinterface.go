// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Rectangle struct {
	length, width int
}

type Shape interface {
	Area2() int
}

func main() {
	var rect1 Rectangle

	rect1.length = 8
	rect1.width = 5
	fmt.Println(Area(rect1.length, rect1.width))
	fmt.Println(rect1.Area2())

	var rect2 Rectangle
	rect2.length = 4
	rect2.width = 6

	var s Shape
	s = &rect2
	fmt.Println(s.Area2())

}

func Area(ln, wd int) int {

	return ln * wd
}

func (x *Rectangle) Area2() int {

	return x.length * x.width
}
