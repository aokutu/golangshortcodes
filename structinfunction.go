// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func main() {

	var r1 rectangle
	r1.length = 5
	r1.width = 4
	fmt.Println(Areax(r1))

}

func Areax(rec rectangle) int {

	return rec.width * rec.length
}
