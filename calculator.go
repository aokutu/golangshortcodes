package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// First number
	fmt.Print("ENTER FIRST NUMBER: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	answer, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	for {
		// Operator
		fmt.Print("ENTER OPERATOR (+, -, *, /, =): ")
		op, _ := reader.ReadString('\n')
		op = strings.TrimSpace(op)

		if op == "=" {
			fmt.Println("Final answer:", answer)
			break
		}

		// Next number
		fmt.Print("ENTER NEXT NUMBER: ")
		numInput, _ := reader.ReadString('\n')
		numInput = strings.TrimSpace(numInput)

		number, err := strconv.Atoi(numInput)
		if err != nil {
			fmt.Println("Invalid number")
			continue
		}

		// Apply operation
		switch op {
		case "+":
			answer += number
		case "-":
			answer -= number
		case "*":
			answer *= number
		case "/":
			if number == 0 {
				fmt.Println("Cannot divide by zero")
				continue
			}
			answer /= number
		default:
			fmt.Println("Invalid operator")
		}
	}
}