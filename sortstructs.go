package main 

import "os"
import "fmt"
import "strings"
import "strconv"
import "sort"

type player struct {
	name string
	scores int
}

func main(){

	players := []player{}
	
	data,_ := os.ReadFile("sample.txt")
	datastring :=  string(data)
	sliceddata:= strings.Split(datastring, "\n")



	for _,details := range sliceddata {

	parts := strings.Fields(details)
    var p player
		p.name = parts[0]
		scr,_ := strconv.Atoi( parts[1] )
		p.scores = scr
		players = append(players,p)

	}

	fmt.Println(players)
	sort.Slice(players, func(i, j int) bool {
        return players[i].scores > players[j].scores
    }   )

	fmt.Println(players)

	rank :=1 
	

	for a:=0 ; a <= len(players) -1  ; a++{
			
		
			  if a > 0 && players[a].scores != players[a-1].scores {
        rank = a + 1 
    }
		fmt.Println( players[a].name ,rank)	
	}

	
} 