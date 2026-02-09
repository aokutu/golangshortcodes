package main 

import "fmt"

func main(){
	fmt.Println(Triarea(6,6) )

}

func Triarea(a,b float64) float64{
	return a *b * 0.5
}