package service

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func Openai() {
	client := openai.NewClient("token")

	resp, err:=client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleUser,
					Content: "hello",
				},
			},
		},
	)
if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
}
fmt.Println(resp.Choices[0].Message.Content)
}