package main

import (
	"fmt"
	"strings"
)

func main() {
	var x strings.Builder
	x.WriteString("ANDREW")
	x.WriteString("OKUTU")
	x.WriteString("OKUTU")
	fmt.Print(x.String())
}