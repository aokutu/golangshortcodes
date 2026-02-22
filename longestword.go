// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	//"sort"
	"unicode"
)

type details struct {
	word     string
	length   int
}

func main() {
	sentence := " 123    456789  &^#@@##&@!!@@^^4  1011"
	words := Wordslice(sentence)

	var detailsslice []details

	pos := 0
	for _, wrd := range words {
		pos++
		word, size := Wordsslicestruct(wrd)
		Wordsslicestruct(wrd)
		var worddetails details
		worddetails.word = word
		worddetails.length = size
		detailsslice = append(detailsslice, worddetails)

	}

	//fmt.Println(detailsslice)

	/*sort.Slice(detailsslice, func(i int, j int) bool {
		return detailsslice[i].length > detailsslice[j].length
	}) */

	Bublesort(detailsslice)
	fmt.Println(detailsslice[0].word)

}

func Wordslice(sentence string) [][]string {

	words := [][]string{}
	sentencerune := []rune(sentence)
	start := 0
	



	for a := 0; a <= len(sentencerune)-1; a++ {

		if sentencerune[a] == ' ' {
			if a > start {
				word := string(sentencerune[start:a])
				wordslice := []string{word}
				words = append(words, wordslice)

			}
			start = a + 1
		}
	}

	if start < len(sentencerune)-1 {
		word := string(sentencerune[start:])
		wordslice := []string{word}
		words = append(words, wordslice)

	}


for pos ,word := range words {

status := Validateword(word)
if status == false {
	words  = append( words[:pos], words[pos+1:]  ...)
}

}
	
	return words
}

func Validateword(word []string) bool {

	wordrunes  := []rune(word[0])

	 

	if (! unicode.IsLetter(wordrunes[0] ) ) &&  (! unicode.IsDigit(wordrunes[0] ) )  {

		return false 
	} 


	for a:= 1; a <  len(word) ; a++ {
		if (! unicode.IsLetter(wordrunes[0] ) ) &&  (! unicode.IsDigit(wordrunes[0] ) )  {

		return false 
	}
	
	if (! unicode.IsLetter(wordrunes[len(word) ] ) ) &&  (! unicode.IsDigit(wordrunes[ len(word) ] ) )  &&   (! unicode.IsPunct(wordrunes[ len(word) ] ) )  {

		return false 
	} 


	}

	
	return true 

}


func Wordsslicestruct(dtl []string) (string, int) {

	word := ""

	for _, ch := range dtl {
		word = word + string(ch)

	}

	return word, len(word)

}




func Bublesort(detailsslice []details) []details {

	for a := 0; a < len(detailsslice)-1; a++ {
		for b := 0; b < len(detailsslice)-1-a; b++ {

			if detailsslice[b].length < detailsslice[b+1].length {
				detailsslice[b], detailsslice[b+1] = detailsslice[b+1], detailsslice[b]

			}

		}

	}

	return detailsslice

}

