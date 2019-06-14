package cmd

import (
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
	git "github.com/todd-bush/go-git-fame/git_client"
	"github.com/todd-bush/go-git-fame/parser"
)

type BlameOutput struct {
	file       string
	blame_data []parser.BlameData
}

type ProcessOutput struct {
	author       string
	email        string
	loc          int
	commits      int
	files        map[string]bool
	file_count   int
	loc_perc     float32
	commits_perc float32
	files_perc   float32
}

type BlameProcess struct {
	file        string
	blame_lines []string
}

func ExecuteProcessor(branch string) []ProcessOutput {

	result := []ProcessOutput{}

	var default_branch string

	if len(branch) > 0 {
		default_branch = branch
	} else {
		default_branch = git.GitCurrentBranch()
	}

	log.Infof("processing fame on branch %s\n", default_branch)

	blame_output := GatherBlame(default_branch)
	commits := GatherCommits()

	log.Infof("commits hash: %v\n", commits)

	for _, blame := range blame_output {

		for _, data := range blame.blame_data {

			if len(data.Mail) == 0 {
				continue
			}

			var author_data *ProcessOutput

			for i := range result {
				if result[i].email == data.Mail {
					author_data = &result[i]
					break
				}
			}

			if author_data == nil {
				var ad = ProcessOutput{
					author:       data.Author,
					email:        data.Mail,
					loc:          0,
					commits:      0,
					files:        make(map[string]bool),
					file_count:   0,
					commits_perc: 0,
					loc_perc:     0,
					files_perc:   0,
				}
				result = append(result, ad)
				author_data = &ad
			}

			log.Infof("about to populate %+v", author_data)

			// add the file
			author_data.files[blame.file] = true

			log.Infof("looking for commit data for %s\n", author_data.email)
			if val, ok := commits[author_data.email]; ok {
				author_data.commits = val
				log.Infof("adding %d to %s\n", val, author_data.email)
			}

			author_data.loc += data.NumLines

		}
	}

	// now do the counts and totals
	total_commits, total_loc, total_files := 0, 0, 0

	for _, out := range result {
		out.file_count = len(out.files)
		total_commits += out.commits
		total_loc += out.loc
		total_files += out.file_count
	}

	log.Infof("totals: %d, %d, %d", total_commits, total_loc, total_files)

	for i, _ := range result {
		var ad *ProcessOutput
		ad = &result[i]

		ad.loc_perc = (float32(ad.loc) / float32(total_loc)) * float32(100)
		ad.commits_perc = float32(ad.commits) / float32(total_commits) * float32(100)
		ad.files_perc = float32(ad.file_count) / float32(total_files) * float32(100)

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
