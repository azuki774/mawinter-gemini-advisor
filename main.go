package main

import (
	"log/slog"
	"os"

	"github.com/azuki774/mawinter-gemini-advisor/cmd"
)

func main() {
	// 1. カスタムハンドラを作成 (例: JSON形式で標準出力へ)
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug, // ログレベルをDebug以上に設定
		AddSource: true,            // ログにソースファイルと行番号を追加
	})
	customLogger := slog.New(jsonHandler)
	slog.SetDefault(customLogger)

	cmd.Execute()
}
