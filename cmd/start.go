/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	// defaultMockValue is the default value for --gemini and --mawinter flags
	defaultMockValue = "mock"
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
		fmt.Printf("start called with --gemini=%s and --mawinter=%s\n", geminiArg, mawinterArg)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringVar(&geminiArg, "gemini", defaultMockValue, "Specify the Gemini related argument")
	startCmd.Flags().StringVar(&mawinterArg, "mawinter", defaultMockValue, "Specify the Mawinter related argument")
}
