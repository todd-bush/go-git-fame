package cmd

import (
	"regexp"
	"strconv"
	"sync"

	pb "github.com/cheggaaa/pb/v3"
	log "github.com/sirupsen/logrus"
	git "github.com/todd-bush/go-git-fame/git_client"
	"github.com/todd-bush/go-git-fame/parser"
)

type BlameOutput struct {
	file      string
	blameData []parser.BlameData
}

type ProcessOutput struct {
	author      string
	email       string
	loc         int
	commits     int
	files       map[string]bool
	fileCount   int
	locPerc     float32
	commitsPerc float32
	filesPerc   float32
}

type BlameProcess struct {
	file        string
	blame_lines []string
}

func ExecuteProcessor(branch string) []ProcessOutput {

	result := []ProcessOutput{}

	var defaultBranch string

	if len(branch) > 0 {
		defaultBranch = branch
	} else {
		defaultBranch = git.GitCurrentBranch()
	}

	log.Infof("processing fame on branch %s\n", defaultBranch)

	blameOutput := GatherBlame(defaultBranch)
	commits := GatherCommits()

	log.Infof("commits hash: %v\n", commits)

	for _, blame := range blameOutput {

		for _, data := range blame.blameData {

			if len(data.Author) == 0 {
				continue
			}

			var authorData *ProcessOutput

			for i := range result {
				if result[i].author == data.Author {
					authorData = &result[i]
					break
				}
			}

			if authorData == nil {
				var ad = ProcessOutput{
					author:      data.Author,
					email:       data.Mail,
					loc:         0,
					commits:     0,
					files:       make(map[string]bool),
					fileCount:   0,
					commitsPerc: 0,
					locPerc:     0,
					filesPerc:   0,
				}
				result = append(result, ad)
				authorData = &ad
			}

			log.Infof("about to populate %+v", authorData)

			// add the file
			authorData.files[blame.file] = true
			authorData.fileCount = len(authorData.files) + 1

			log.Infof("looking for commit data for %s\n", authorData.author)
			if val, ok := commits[authorData.author]; ok {
				authorData.commits = val
				log.Infof("adding %d to %s\n", val, authorData.author)
			}

			authorData.loc += data.NumLines

		}
	}

	// now do the counts and totals
	totalCommits, totalLoc, totalFiles := 0, 0, 0

	for _, out := range result {
		out.fileCount = len(out.files)
		totalCommits += out.commits
		totalLoc += out.loc
		totalFiles += out.fileCount
	}

	log.Infof("totals: %d, %d, %d", totalCommits, totalLoc, totalFiles)

	for i, _ := range result {
		var ad *ProcessOutput
		ad = &result[i]

		ad.locPerc = (float32(ad.loc) / float32(totalLoc)) * float32(100)
		ad.commitsPerc = float32(ad.commits) / float32(totalCommits) * float32(100)
		ad.filesPerc = float32(ad.fileCount) / float32(totalFiles) * float32(100)

	}

	return result

}

func GatherBlame(branch string) []BlameOutput {

	// get this list of files
	fileList := git.GitListFiles(branch)

	log.Infof("found %d files to process", len(fileList))

	blameOut := []BlameProcess{}

	bar := pb.StartNew(len(fileList))

	// default to 5 routines
	guard := make(chan struct{}, 5)
	var wg sync.WaitGroup

	wg.Add(len(fileList))

	for _, file := range fileList {

		bar.Increment()
		guard <- struct{}{}

		go func(f string) {

			defer wg.Done()

			if len(f) > 0 {

				blameResult := git.GitBlame(f)
				out := BlameProcess{
					file:        f,
					blame_lines: blameResult,
				}

				blameOut = append(blameOut, out)
			}
			<-guard
		}(file)
	}

	wg.Wait()

	blameCollector := []BlameOutput{}

	for _, bi := range blameOut {

		log.Infof("parsing blame on file: %s; lines %d", bi.file, len(bi.blame_lines))

		if len(bi.blame_lines) > 1 {
			blameOut := parser.Parse(bi.blame_lines)
			blameCollector = append(blameCollector, BlameOutput{
				file:      bi.file,
				blameData: blameOut,
			})
		}
	}

	bar.Finish()

	return blameCollector
}

func GatherCommits() map[string]int {
	result := map[string]int{}

	commitLines := git.GitShortLog()

	r, _ := regexp.Compile(`(\d+)\s+(.+)\s+<(.+?)>`)

	for _, commitLine := range commitLines {

		if len(commitLine) > 0 {

			log.Debugf("parsing line %s", commitLine)
			peices := r.FindStringSubmatch(commitLine)

			log.Debugf("peices = %v", peices)
			commits, _ := strconv.Atoi(peices[1])
			result[peices[2]] = commits
		}
	}

	return result
}
