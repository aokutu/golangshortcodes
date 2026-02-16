package main 

import "fmt" 
import "os"

func main(){
	olddata,_:= os.ReadFile("data.txt")

	olddata = append(olddata,[]byte("\nabc") ...)

	os.WriteFile("data.txt",olddata,0644)
		
fmt.Println(olddata)

olddata,_= os.ReadFile("data.txt")


}