package main

import (
	"bufio"
	"fmt"
	"os"
)

import "aihackerthon/agent"


func main() {
	ai, err := agent.NewAgent()
	if err != nil {
		fmt.Println("", err)
		return
	}

	fmt.Println(" AI Agent ready. Type 'exit' to quit.")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nYou: ")
		scanner.Scan()
		text := scanner.Text()

		if text == "exit" {
			break
		}

		reply, err := ai.Ask(text)
		if err != nil {
			fmt.Println("", err)
			continue
		}

		fmt.Println("🤖:", reply)
	}
}
