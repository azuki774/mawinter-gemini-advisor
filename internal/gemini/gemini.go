package gemini

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"golang.org/x/exp/slog"
	"google.golang.org/api/option"
)

type GeminiClient struct {
	UseModel       string
	GeminiAPIKey   string
	CustomEndpoint string
}

func NewGeminiClient(useModel string, geminiAPIKey string, customEndpoint string) *GeminiClient {
	if customEndpoint == "mock" {
		customEndpoint = "http://localhost:8000"
	}

	return &GeminiClient{
		UseModel:       useModel,
		GeminiAPIKey:   geminiAPIKey,
		CustomEndpoint: customEndpoint,
	}
}

func (g *GeminiClient) Post(ctx context.Context, prompt string) (string, error) {
	var client *genai.Client
	var err error

	if g.CustomEndpoint != "" {
		slog.Info("custom endpoint is used")
		client, err = genai.NewClient(ctx, option.WithAPIKey(g.GeminiAPIKey), option.WithEndpoint(g.CustomEndpoint))
		if err != nil {
			slog.Error("gemini call error", "err", err)
			return "", err
		}
		defer client.Close()
	} else {
		client, err = genai.NewClient(ctx, option.WithAPIKey(g.GeminiAPIKey))
		if err != nil {
			slog.Error("gemini call error", "err", err)
			return "", err
		}
		defer client.Close()
	}

	// 利用するモデルを指定します (例: "gemini-2.5-flash-preview-04-17")
	model := client.GenerativeModel(g.UseModel)

	slog.Info("gemini call", "model", g.UseModel)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		slog.Error("GenerateContent failed", "err", err)
	}

	if resp != nil && len(resp.Candidates) > 0 {
		for _, cand := range resp.Candidates {
			if cand.Content != nil {
				for _, part := range cand.Content.Parts {
					if text, ok := part.(genai.Text); ok {
						slog.Info("gemini response")
						return string(text), nil
					}
				}
			}
		}
	} else {
		return "", fmt.Errorf("no response found")
	}
	return "", fmt.Errorf("unknown error")
}

func init() {
	// 1. カスタムハンドラを作成 (例: JSON形式で標準出力へ)
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug, // ログレベルをDebug以上に設定
		AddSource: true,            // ログにソースファイルと行番号を追加
	})
	customLogger := slog.New(jsonHandler)
	slog.SetDefault(customLogger)
}
