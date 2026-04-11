package main 

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"
import "sort"


type playerdetails struct {
	player string
	score int
}

func main(){

	plyrs := []playerdetails{}

	for {

		fmt.Print("ENTER PLAYER:")
	 
	

	reader := bufio.NewReader( os.Stdin)
	plx,err := reader.ReadString('\n')
	player := strings.TrimSpace(plx)
	if player == "EXIT"{
		break
	}
	 
	if err != nil {

	}
     
	fmt.Print("ENTER SCORES:")

	scores,_ := reader.ReadString('\n')
	
	scoreint,_ := strconv.Atoi( strings.TrimSpace(scores) )


	

	addplyr := addplayers(player, scoreint) 
	
	plyrs =append(plyrs,addplyr)
	
	}

	
fmt.Println(plyrs)
sort.Slice(plyrs, func(i int, j int) bool {
	return plyrs[i].score > plyrs[j].score 
}  )

	//fmt.Print(plyrs )
	rnk :=1	
	//rnk2 :=1 
	for a:=0 ; a < len(plyrs)-1  ; a++{
	    if a > 0 && plyrs[a].score < plyrs[a-1].score {
        rnk = a + 1
    }
    fmt.Println(plyrs[a].player, rnk)

	}
	

} 

func addplayers(player string,scores int)  playerdetails {

	var  pl playerdetails
	pl.player = player
	pl.score = scores





	return pl

}