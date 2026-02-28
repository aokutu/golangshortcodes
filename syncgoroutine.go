package main 

import "fmt" 
import "sync"

func main(){
	
	var wt  sync.WaitGroup

	wt.Add(2)

	go func (){
		defer wt.Done()
		for a:='A' ; a <= 'Z' ;a++{
			fmt.Print( string(a))
			
		}
	} ()


		go  func (){
			defer wt.Done()
		for a:='a' ; a <= 'z' ;a++{
			fmt.Print( string(a))
			
		}
	} ()

  wt.Wait()
}