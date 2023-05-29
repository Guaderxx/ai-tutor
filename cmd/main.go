package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/Guaderxx/ai-english-tutor/pkg/audio"
	"github.com/Guaderxx/ai-english-tutor/pkg/runpy"
	"github.com/sashabaranov/go-openai"
)

func main() {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "你是一个有善的英语外教，会和我进行日常对话，并纠正我的语法错误",
			},
		},
		// Stream: true,   // TODO
	}

	fmt.Println("Conversation")
	fmt.Println("---------------------")
	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)

	wg := &sync.WaitGroup{}

	for s.Scan() {
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: s.Text(),
		})
		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		content := resp.Choices[0].Message.Content
		// convert to audio before print
		audioPath := runpy.TextToAudio(content)

		wg.Add(1)

		go func() {
			audio.PlayAudio(audioPath)
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			fmt.Printf("gpt: ")
			for _, c := range content {
				fmt.Printf("%c", c)
				time.Sleep(12 * time.Millisecond)
			}
			fmt.Println("\n----------------")
			wg.Done()
		}()

		wg.Wait()

		req.Messages = append(req.Messages, resp.Choices[0].Message)
		fmt.Print("> ")
	}
}
