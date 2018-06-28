package main

import (
	"testing"
)

func Test_git_list_files(t *testing.T) {
	files := git_list_files(".", "develop")

	if len(files) <= 0 {
		t.Fatal("Expected files to be larger than zero")
	}
}
