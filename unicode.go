package main

import (
    "fmt"
    "unicode"
)

func main() {
    tests := []rune{'A', 'a', '9', 'é', '汉', ' ', '\n', 'π', '©'}

    for _, r := range tests {
        fmt.Printf("Rune: %c (U+%04X)\n", r, r)
        fmt.Printf("  IsLetter:  %v\n", unicode.IsLetter(r))
        fmt.Printf("  IsUpper:   %v\n", unicode.IsUpper(r))
        fmt.Printf("  IsLower:   %v\n", unicode.IsLower(r))
        fmt.Printf("  IsDigit:   %v\n", unicode.IsDigit(r))
        fmt.Printf("  IsSpace:   %v\n", unicode.IsSpace(r))
        fmt.Printf("  IsPrint:   %v\n", unicode.IsPrint(r))
        fmt.Printf("  In Latin:  %v\n", unicode.Is(unicode.Latin, r))
        fmt.Printf("  In Han:    %v\n", unicode.Is(unicode.Han, r))
		 fmt.Printf("  In Han:    %v\n", unicode.IsPunct('r'))
         /*
         unicode.IsLetter(r)     // true - letters (L category)
unicode.IsDigit(r)      // false - decimal digits (Nd)
unicode.IsNumber(r)     // false - any number (N)
unicode.IsPunct(r)      // false - punctuation (P)
unicode.IsSymbol(r)     // false - symbols (S)
unicode.IsSpace(r)      // false - whitespace (Z)
unicode.IsGraphic(r)    // true - visible glyphs (L,M,N,P,S,Zs)
unicode.IsControl(r)    // false - control characters (C)
unicode.IsPrint(r)      // true - printable characters
unicode.IsUpper(r)      // true - uppercase
unicode.IsLower(r)      // false - lowercase
         */

        fmt.Println("---")
    }
}