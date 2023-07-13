package utilities

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func GetData(text string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error in loading .env file")
	}
	key := os.Getenv("API_KEY")
	//fmt.Println(key)

	//text := "Who is PM of India"
	c := openai.NewClient(key)

	resp, err := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: text,
				},
			},
		},
	)

	if err != nil {
		panic(err)

	}
	te := resp.Choices[0].Message.Content
	fmt.Println(resp.Choices[0].Message.Content)
	fmt.Println()
	return te
}
