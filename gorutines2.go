// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

func main() {
	go Test1()
	time.Sleep(2 * time.Second)
	fmt.Println()
	go Test2()
	time.Sleep(2 * time.Second)
}

func Test1() {

	for a := 'a'; a <= 'z'; a++ {
		fmt.Print(string(a))
	}
}

func Test2() {

	for a := 'A'; a <= 'Z'; a++ {
		fmt.Print(string(a))
	}
}
