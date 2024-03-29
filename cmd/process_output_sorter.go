package cmd

import "sort"

type By func(p1, p2 *ProcessOutput) bool

type processOutputSorter struct {
	processOutput []ProcessOutput
	by            By
}

func (by By) Sort(po []ProcessOutput) {
	ps := &processOutputSorter{
		processOutput: po,
		by:            by,
	}
	sort.Sort(ps)
}

// Len is part of the sort.Interface
func (s *processOutputSorter) Len() int {
	return len(s.processOutput)
}

// Swap is part of the sort.Interface
func (s *processOutputSorter) Swap(i, j int) {
	s.processOutput[i], s.processOutput[j] = s.processOutput[j], s.processOutput[i]
}

func (s *processOutputSorter) Less(i, j int) bool {
	return s.by(&s.processOutput[i], &s.processOutput[j])
}

// ByCommits sorts by number of commits
func ByCommits(p1, p2 *ProcessOutput) bool {
	return p1.commits > p2.commits
}

// ByLoc sorts by number of LOC
func ByLoc(p1, p2 *ProcessOutput) bool {
	return p1.loc > p2.loc
}

// ByFiles sorts by files
func ByFiles(p1, p2 *ProcessOutput) bool {
	return p1.fileCount > p2.fileCount
}
