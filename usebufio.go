package main 

import "fmt"
import "bufio"
import "os"
import "strings"

func main () {

	for {

reader:=bufio.NewReader(os.Stdin)

	fmt.Print("ENTER NAME ")

	name,_:=reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("ENTER SECOND NAME  ")

	name2,_:=reader.ReadString('\n')
	name2 = strings.TrimSpace(name2)

	fmt.Print("TERMINATE ? (Y/N) ")

	terminate,_:=reader.ReadString('\n')
	terminate = strings.TrimSpace(terminate)

	if terminate  == "Y" {
		fmt.Println(name,name2)
		break

	} else if terminate == "N" {

fmt.Println(name,name2)
		continue
	}



	}

	
}