package main 

import "fmt" 
import "sync"

func main(){

	var wg sync.WaitGroup

	wg.Add(4)

go Test1(&wg) 
go Test2(&wg)
go Test3(&wg)
go Test4("END",&wg)

wg.Wait()

}

func Test1(wg *sync.WaitGroup){
	defer wg.Done()

	a :=1 
	for a <= 100 {
		fmt.Println(a)
		a+=1
	}
}

func Test2(wg *sync.WaitGroup){
	defer wg.Done()

	a :=100 
	for a <= 200 {
		fmt.Println(a)
		a+=1
	}

	
}

func Test3(wg *sync.WaitGroup){
	defer wg.Done()

	a :=200 
	for a <= 300 {
		fmt.Println(a)
		a+=1
	}
}

func Test4(end string,wg *sync.WaitGroup){
	defer wg.Done()

		
	fmt.Println(end)

}