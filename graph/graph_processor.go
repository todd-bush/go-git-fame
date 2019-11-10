package graph

import (
	"sort"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	git "github.com/todd-bush/go-git-fame/git_client"
)

var gitDateForm = "Mon Jan 2 15:04:05 2006 -0700"

func CollectCommits() []CommitsByDate {
	commits := git.GitCommitDate()
	dateMap := make(map[time.Time]float64)
	dayDuration, _ := time.ParseDuration("24h")

	for _, c := range commits {

		dateString := strings.Split(c, "::")[1]

		commitDate, err := time.Parse(gitDateForm, dateString)

		if err != nil {
			log.Errorf("Error parsing date %s.  Error: %v", dateString, err)
		}

		commitDate = commitDate.Truncate(dayDuration)

		if val, ok := dateMap[commitDate]; ok {
			dateMap[commitDate] = val + 1
			log.Infof("Updating %+v, value %f", commitDate, dateMap[commitDate])
		} else {
			log.Infof("Adding %+v to map", commitDate)
			dateMap[commitDate] = 1
		}
	}

	commitDates := make([]CommitsByDate, len(dateMap))

	i := 0
	for k, v := range dateMap {
		commitDates[i] = CommitsByDate{date: k, commits: v}
		i++
	}

	sort.SliceStable(commitDates, func(i, j int) bool {
		return commitDates[i].date.After(commitDates[j].date)
	})

	return commitDates
}

func GraphCommits(cbd []CommitsByDate) {
	err := GraphCommitsByTime(cbd)

	if err != nil {
		log.Fatalf("error to create image file: %+v", err)
	}
}
