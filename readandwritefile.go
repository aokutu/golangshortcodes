package main

import (
	
	"fmt"
	"os"
)

func main() {
	///IT  ACTUALLY APPENDS  NEW  LINE 
	oldcontent,_:= os.ReadFile("data.txt")

	newstring := os.Args[1] + "\n"
	
	newdata := []byte( newstring  )

	newdata =append(oldcontent, newdata ...)

	os.WriteFile("data.txt",newdata,0644)
	content,_:= os.ReadFile("data.txt")
	
	fmt.Println( string(content))

    

	

}