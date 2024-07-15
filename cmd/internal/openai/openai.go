package openai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type Config struct {
	APIKey string
}

type Client struct {
	apiClient *openai.Client
}

func NewClient(cfg Config) *Client {
	apiClient := openai.NewClient(cfg.APIKey)
	return &Client{
		apiClient: apiClient,
	}
}

func (cli *Client) MakeRequest(requestStr string) (string, error) {
	resp, err := cli.apiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o20240513,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: requestStr,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("OPENAI - MakeRequest error: %v", err)
	} else {
		return resp.Choices[0].Message.Content, nil
	}
}
