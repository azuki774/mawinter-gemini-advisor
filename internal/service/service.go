package service

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"time"
)

type Service struct {
	GeminiClient   GeminiClient
	MawinterClient MawinterClient
	FileOperator   FileOperator
	PrePromptFile  string
	OutputDir      string
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

func NewService(geminiClient GeminiClient, mawinterClient MawinterClient, fileOperator FileOperator, prePromptFile string, outputDir string) *Service {
	return &Service{
		GeminiClient:   geminiClient,
		MawinterClient: mawinterClient,
		FileOperator:   fileOperator,
		PrePromptFile:  prePromptFile,
		OutputDir:      outputDir,
	}
}

func buildPrompt(prePrompt string, mawinterStr string) string {
	return mawinterExplainPrompt + "--------------\n" + prePrompt + "--------------\n" + mawinterStr
}

func (s *Service) Start(ctx context.Context) error {
	slog.Info("start service")
	// pre_prompt ファイルを読み込む
	prePrompt, err := s.FileOperator.LoadTxtFile(s.PrePromptFile)
	if err != nil {
		slog.Error("failed to load pre-prompt file", "filename", s.PrePromptFile, "error", err)
		return err
	}
	slog.Info("pre-prompt file loaded", "filename", s.PrePromptFile)

	// mawinter API を呼び出して、今月のデータを取得する
	mawinterStr, err := s.MawinterClient.GetMonthlyData(time.Now().Format("200601"))
	if err != nil {
		slog.Error("failed to get monthly data from mawinter", "error", err)
		return err
	}

	// mawinter API を呼び出して、先月のデータを取得する
	mawinterLastMonthStr, err := s.MawinterClient.GetMonthlyData(time.Now().AddDate(0, -1, 0).Format("200601"))
	if err != nil {
		slog.Error("failed to get monthly data from mawinter (lastmonth)", "error", err)
		return err
	}

	prompt := buildPrompt(prePrompt, mawinterStr+"--------------\n"+mawinterLastMonthStr)

	// gemini API を呼び出して、レスポンスを取得する
	advRes, err := s.GeminiClient.Post(ctx, prompt)
	if err != nil {
		slog.Error("failed to get response from gemini", "error", err)
		return err
	}

	// gemini API のレスポンスをファイルに書き出す
	responseFileName := fmt.Sprintf("gemini_%s.txt", time.Now().Format("20060102"))
	err = s.FileOperator.WriteTxtFile(filepath.Join(s.OutputDir, responseFileName), advRes)
	if err != nil {
		slog.Error("failed to write response from gemini", "error", err)
		return err
	}

	slog.Info("advise gemini response written")
	return nil
}
