package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	branch string
)

var rootCmd = &cobra.Command{
	Use:   "go-git-fame",
	Short: "Fame give you commit stats for your GIT repo",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&branch, "branch", "b", "branch to use, defaults to current HEAD")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
