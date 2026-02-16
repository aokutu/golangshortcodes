package main

import "fmt"
import "bufio"
import "os"

func main(){

	readfile,_ :=os.Open("data.txt")

	reader := bufio.NewScanner(readfile)

	for reader.Scan(){

	fmt.Println( reader.Text())

	}



}