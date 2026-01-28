package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var x string
	a := "ANDREW"
	b, _ := json.Marshal(a)
	c := string(b)
	fmt.Println(c)
	json.Unmarshal(b, &x)

	fmt.Println(x)

}