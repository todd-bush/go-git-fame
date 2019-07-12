package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	branch  string
	verbose bool
	s       string
)

var rootCmd = &cobra.Command{
	Use:   "go-git-fame",
	Short: "Fame give you commit stats for your GIT repo",
	Run: func(cmd *cobra.Command, args []string) {

		if verbose {
			log.SetLevel(log.InfoLevel)
		} else {
			log.SetLevel(log.ErrorLevel)
		}

		output := ExecuteProcessor(branch)

		// default
		if s == "loc" {
			By(ByLoc).Sort(output)
		} else if s == "files" {
			By(ByFiles).Sort(output)
		} else {
			By(ByCommits).Sort(output)
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Author", "Files", "Commits", "LOC", "Distribution"})

		for _, out := range output {
			dist := fmt.Sprintf("%04.2f/%04.2f/%04.2f", out.filesPerc, out.commitsPerc, out.locPerc)
			t.AppendRow(table.Row{out.author, out.fileCount, out.commits, out.loc, dist})
		}

		t.Render()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&branch, "branch", "", "branch to use, defaults to current HEAD")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "enable verbosness")
	rootCmd.PersistentFlags().StringVar(&s, "sort", "", "sort field, either 'commit' (default), 'loc', 'files'")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
