package API

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func NewApiCall(apiKey string, prompt string) string {

	// openai.Client instance
	client := openai.NewClient(apiKey)
	context := context.Background()

	// GPT-3 model
	reqst := openai.CompletionRequest{
		Model: openai.GPT3TextDavinci003, MaxTokens: 1080, Prompt: prompt,
	}

	// Completion request to the OpenAI API
	resp, err := client.CreateCompletion(context, reqst)
	if err != nil {
		return fmt.Sprintf("%d", err)
	}

	// Return API response text
	return resp.Choices[0].Text
}
