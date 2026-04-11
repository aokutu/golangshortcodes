// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	x := 5
	fmt.Println(x)
	changex(x)
	fmt.Println(x)
	changey(&x)
	fmt.Println(x)

}

func changex(x int) {
	x = 6

}

func changey(x *int) {

	*x = 6
}
