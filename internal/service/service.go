package service

import (
	"context"
	"log/slog"
	"time"
)

type Service struct {
	GeminiClient   GeminiClient
	MawinterClient MawinterClient
	FileOperator   FileOperator
}

type GeminiClient interface {
	Post(ctx context.Context, prompt string) (string, error)
}

type MawinterClient interface {
	GetMonthlyData(yyyymm string) (string, error) // response JSON string
}

type FileOperator interface {
	LoadTxtFile(filePath string) (string, error)
	WriteTxtFile(filePath string, data string) error
}

func NewService(geminiClient GeminiClient, mawinterClient MawinterClient, fileOperator FileOperator) *Service {
	return &Service{
		GeminiClient:   geminiClient,
		MawinterClient: mawinterClient,
		FileOperator:   fileOperator,
	}
}

func buildPrompt(prePrompt string, mawinterStr string) string {
	return prePrompt + "--------------\n" + mawinterStr
}

func (s *Service) Start(ctx context.Context) error {
	slog.Info("start service")
	// pre_prompt ファイルを読み込む
	prePrompt, err := s.FileOperator.LoadTxtFile("pre_prompt.txt")
	if err != nil {
		slog.Error("failed to load pre_prompt.txt", "error", err)
		return err
	}
	slog.Info("pre_prompt.txt loaded")

	// mawinter API を呼び出して、今月のデータを取得する
	mawinterStr, err := s.MawinterClient.GetMonthlyData(time.Now().Format("200601"))
	if err != nil {
		slog.Error("failed to get monthly data from mawinter", "error", err)
		return err
	}

	prompt := buildPrompt(prePrompt, mawinterStr)

	// gemini API を呼び出して、レスポンスを取得する
	advRes, err := s.GeminiClient.Post(ctx, prompt)
	if err != nil {
		slog.Error("failed to get response from gemini", "error", err)
		return err
	}

	// gemini API のレスポンスをファイルに書き出す
	err = s.FileOperator.WriteTxtFile("adv_res.txt", advRes)
	if err != nil {
		slog.Error("failed to write response from gemini", "error", err)
		return err
	}

	slog.Info("advise gemini response written")
	return nil
}
