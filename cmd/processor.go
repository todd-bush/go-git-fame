package cmd

import (
	log "github.com/sirupsen/logrus"
	git "github.com/todd-bush/go-git-fame/git_client"
	"github.com/todd-bush/go-git-fame/parser"
	"regexp"
	"strconv"
)

type BlameOutput struct {
	file       string
	blame_data []parser.BlameData
}

type ProcessOutput struct {
	author       string
	loc          int
	commits      int
	files        int
	loc_perc     float32
	commits_perc float32
	files_perc   float32
}

type BlameProcess struct {
	file        string
	blame_lines []string
}

func ExecuteProcessor() []ProcessOutput {

	result := []ProcessOutput{}

	blame_output := GatherBlame("master") // TODO should be an argument
	commits := GatherCommits()

	for _, blame := range blame_output {
		for _, data := range blame.blame_data {

			author_data := ProcessOutput{}

			for i := range result {
				if result[i].author == data.Mail {
					author_data = result[i]
					break
				}
			}

			if len(author_data.author) == 0 {
				author_data := ProcessOutput{
					author: data.Mail,
					loc:    0,
				}
				result = append(result, author_data)
			}

			com := 0

			if val, ok := commits[data.Mail]; ok {
				com = val
			}

			author_data.loc += data.Num_lines
			author_data.commits = com
		}
	}

	return result

}

func GatherBlame(branch string) []BlameOutput {

	// get this list of files
	file_list := git.GitListFiles(branch)

	log.Infof("found %d files to process", len(file_list))

	blame_out := []BlameProcess{}

	for _, file := range file_list {

		if len(file) > 0 {

			blame_result := git.GitBlame(file)
			out := BlameProcess{
				file:        file,
				blame_lines: blame_result,
			}

			blame_out = append(blame_out, out)
		}
	}

	blame_collector := []BlameOutput{}

	for _, bi := range blame_out {

		log.Infof("parsing blame on file: %s", bi.file)

		blame_out := parser.Parse(bi.blame_lines)
		blame_collector = append(blame_collector, BlameOutput{
			file:       bi.file,
			blame_data: blame_out,
		})
	}

	return blame_collector
}

func GatherCommits() map[string]int {
	result := map[string]int{}

	commit_lines := git.GitShortLog()

	r, _ := regexp.Compile(`(\d+)\s+(.+)\s+<(.+?)>`)

	for _, commit_line := range commit_lines {

		if len(commit_line) > 0 {

			log.Debugf("parsing line %s", commit_line)
			peices := r.FindStringSubmatch(commit_line)

			log.Debugf("peices = %v", peices)
			commits, _ := strconv.Atoi(peices[1])
			result[peices[3]] = commits
		}
	}

	return result
}
