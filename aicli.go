package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const openRouterURL = "https://openrouter.ai/api/v1/chat/completions"

func main() {
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		fmt.Println("OPENROUTER_API_KEY not set")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("ENTER QUESTION: ")
		question, _ := reader.ReadString('\n')
		question = strings.TrimSpace(question)

		

		if question == "" {
			fmt.Println("Please enter a question.")
			continue
		}

		// Ask OpenRouter and get only the AI answer
		answer, err := askOpenRouter(apiKey, question)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("\nAnswer:\n")
			fmt.Println(answer) // only the AI response
		}


		fmt.Print("TERMINATE? (Y/N): ")
		terminate, _ := reader.ReadString('\n')
		terminate = strings.TrimSpace(strings.ToUpper(terminate))
		if terminate == "Y" || terminate == "y"  {
			break
		}
	}
}

func askOpenRouter(apiKey, question string) (string, error) {
	payload := map[string]interface{}{
		"model": "gpt-4o-mini",
		"messages": []map[string]string{
			{"role": "user", "content": question},
		},
	}
	data, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", openRouterURL, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	// check for errors returned by API
	if errObj, ok := result["error"].(map[string]interface{}); ok {
		return "", fmt.Errorf("%v", errObj["message"])
	}

	// extract only the content from the first choice
	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("no choices returned")
	}

	message, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("no message returned")
	}

	content, ok := message["content"].(string)
	if !ok {
		return "", fmt.Errorf("no content in message")
	}

	return content, nil
}
