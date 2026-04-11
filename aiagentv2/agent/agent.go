package agent

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Agent struct {
	Messages []Message
	ApiKey   string
}

func NewAgent() (*Agent, error) {
	key := os.Getenv("OPENROUTER_API_KEY")
	if key == "" {
		return nil, errors.New("OPENROUTER_API_KEY not set")
	}

return &Agent{
	ApiKey: key,
	Messages: []Message{
		{
			Role: "system",
			Content: `
You are a Kenyan CBC (Competency-Based Curriculum) learning assistant.

STRICT RULES:
- You ONLY teach Kenyan CBC curriculum.
- You ONLY support Grade 6, Grade 7, Grade 8, and Grade 9 learners.
- Use simple, child-friendly language.
- Explain step by step like a teacher.
- Give examples relevant to Kenya.

SUBJECTS YOU CAN TEACH:
- Mathematics
- English
- Kiswahili
- Integrated Science
- Social Studies
- CRE
- Pre-Technical Studies
- Agriculture
- Computer Studies

REFUSAL RULES:
- If asked about university, college, or adult topics, politely refuse.
- If asked about non-CBC systems (IGCSE, IB, GCSE, A-Levels), refuse.
- If asked topics beyond Grade 9, refuse.

When refusing, say:
"I'm here to help Kenyan CBC students from Grade 6 to Grade 9 only."

If the learner does not mention their grade, ask them which grade (6–9) they are in.
`,
		},
	},
}, nil
}

func (a *Agent) Ask(userInput string) (string, error) {
	// Add user message
	a.Messages = append(a.Messages, Message{
		Role:    "user",
		Content: userInput,
	})

	body := map[string]interface{}{
		"model":    "mistralai/mistral-7b-instruct",
		"messages": a.Messages,
	}

	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest(
		"POST",
		"https://openrouter.ai/api/v1/chat/completions",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+a.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("HTTP-Referer", "http://localhost")
	req.Header.Set("X-Title", "Go AI Agent")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal(raw, &result)

	if errObj, ok := result["error"]; ok {
		return "", errors.New(errObj.(map[string]interface{})["message"].(string))
	}

	choices := result["choices"].([]interface{})
	message := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	reply := message["content"].(string)

	// Save assistant reply
	a.Messages = append(a.Messages, Message{
		Role:    "assistant",
		Content: reply,
	})

	return reply, nil
}
