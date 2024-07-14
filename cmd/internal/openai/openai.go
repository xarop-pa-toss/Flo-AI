package openai

import (
	"context"
	"errors"
	"fmt"

	"Flo-AI/cmd/internal/viper"

	"github.com/sashabaranov/go-openai"
)

func LoadConfig() error {
	token, err := viper.GetString("OPENAI_TOKEN")
	if err != nil {
		return errors.New("OPENAI - No response from Viper for token")
	}

	client := *openai.NewClient((token))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o20240513,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello.",
				},
			},
		},
	)
	if err != nil {
		return fmt.Errorf("Chat Completion error: %v\n", err)
	}

	fmt.Println(resp.Choices[0].Message.Content)
	return nil
}
