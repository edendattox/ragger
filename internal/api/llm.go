package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	constant "github.com/edendattox/ragger/constants"
)

type OpenAI struct {
	APIKey string
}

type Message struct {
	Role    string
	Content string
}

type ChatCompletionRequest struct {
	Messages []Message `json:"messages"`
	Model    string    `json:"model"`
	Stream   bool      `json:"stream"`
}

type ChatCompletionResponse struct {
	Choices []struct {
		Text     string `json:"text"`
		Index    int    `json:"index"`
		Finish   string `json:"finish"`
		To       string `json:"to"`
		Selected bool   `json:"selected"`
		Label    string `json:"label"`
	} `json:"choices"`
}

func (o *OpenAI) GetChatCompletions(model string, message []Message, stream bool) (*ChatCompletionResponse, error) {
	// Prepare request body
	requestBody := &ChatCompletionRequest{
		Messages: message,
		Model:    model,
		Stream:   stream,
	}

	// Marshal request body to JSON
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", constant.OPENAI_URL, bytes.NewBuffer(requestBodyBytes))

	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

	// Set authorization header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+o.APIKey)

	// Send HTTP request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Decode response body
	var response ChatCompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return &response, nil
}
