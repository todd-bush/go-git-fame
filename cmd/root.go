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
	debug   bool
)

var rootCmd = &cobra.Command{
	Use:   "go-git-fame",
	Short: "Fame give you commit stats for your GIT repo",
	Run: func(cmd *cobra.Command, args []string) {

		if debug {
			log.SetLevel(log.DebugLevel)
		} else if verbose {
			log.SetLevel(log.InfoLevel)
		} else {
			log.SetLevel(log.ErrorLevel)
		}

		output := ExecuteProcessor(branch)

		// default
		switch s {
		case "loc":
			By(ByLoc).Sort(output)
		case "files":
			By(ByFiles).Sort(output)
		default:
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

var graphCommand = &cobra.Command{
	Use:   "version",
	Short: "Creates and persists a PNG graph of commits over time",
	Long:  "Creates and persists a PNG graph image of commits over time",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&branch, "branch", "", "branch to use, defaults to current HEAD")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "enable verbosness")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug")
	rootCmd.PersistentFlags().StringVar(&s, "sort", "", "sort field, either 'commit' (default), 'loc', 'files'")

	rootCmd.AddCommand(graphCommand)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
