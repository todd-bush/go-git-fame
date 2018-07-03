package cmd

import (
	"github.com/todd-bush/go-git-fame/parser"
)

type BlameInput struct {
	file       string
	blame_data parser.BlameData
}

type BlameOutput struct {
	author       string
	loc          int32
	commits      int32
	files        int32
	loc_perc     float32
	commits_perc float32
	files_perc   float32
}

func ProcessBlame(inputs []BlameInput) []BlameOutput {

	result := []BlameOutput{}

	return result
}
