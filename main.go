package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/azuki774/mawinter-gemini-advisor/cmd"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	cmd.Execute()

	ctx := context.Background()
	// APIキーを環境変数から取得するか、直接指定します。
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable not set.")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 利用するモデルを指定します (例: "gemini-pro")
	model := client.GenerativeModel("gemini-2.5-flash-preview-04-17") // または "gemini-pro" など、利用可能なモデル名

	prompt := genai.Text("Go言語でGemini APIを使うためのサンプルコードを教えてください。")

	resp, err := model.GenerateContent(ctx, prompt)
	if err != nil {
		log.Fatalf("GenerateContent failed: %v", err)
	}

	if resp != nil && len(resp.Candidates) > 0 {
		for _, cand := range resp.Candidates {
			if cand.Content != nil {
				for _, part := range cand.Content.Parts {
					if text, ok := part.(genai.Text); ok {
						fmt.Println(text)
					}
				}
			}
		}
	} else {
		fmt.Println("No response content found.")
	}
}
