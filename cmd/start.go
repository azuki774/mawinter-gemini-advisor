package cmd

import (
	"log/slog"
	"os"

	"github.com/azuki774/mawinter-gemini-advisor/internal/fileoperator"
	"github.com/azuki774/mawinter-gemini-advisor/internal/gemini"
	"github.com/azuki774/mawinter-gemini-advisor/internal/mawinter"
	"github.com/azuki774/mawinter-gemini-advisor/internal/service"
	"github.com/spf13/cobra"
)

const (
	// defaultMockValue is the default value for --gemini and --mawinter flags
	defaultMockValue = "mock"
	useGeminiModel   = "gemini-2.5-flash-preview-04-17"
)

var (
	geminiArg   string
	mawinterArg string
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("start called with values", "geminiArg", geminiArg, "mawinterArg", mawinterArg)
		mc := mawinter.NewMawinterClient(mawinterArg)
		gc := gemini.NewGeminiClient(useGeminiModel, os.Getenv("GEMINI_API_KEY"), geminiArg)
		fc := fileoperator.NewFileOperator()
		service.NewService(gc, mc, fc).Start(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringVar(&geminiArg, "gemini", "", "Specify the Gemini related argument") // 空文字の場合は本物に接続
	startCmd.Flags().StringVar(&mawinterArg, "mawinter", defaultMockValue, "Specify the Mawinter related argument")
}
