package main 

import "fmt"

func main(){

ch1 := make(chan int)

 go Testchan(ch1)

 y := <- ch1 

 fmt.Println(y)

 fmt.Println()

 go Testchan(ch1)

x := <-ch1


fmt.Println(x)

}

func  Testchan(ch1 chan int){
	

	for a:= 0 ; a <= 50 ; a++ {
		fmt.Print(a)
	}
	ch1 <- 5 

}