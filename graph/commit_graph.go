package graph

import (
	"bytes"
	"io/ioutil"
	"time"

	chart "github.com/wcharczuk/go-chart"
)

type CommitsByDate struct {
	date    time.Time
	commits float64
}

var CommitByDateFileName = "CommitsByTime.png"

func GraphCommitsByTime(commits []CommitsByDate) error {

	values := make([]chart.Value, len(commits))

	for _, c := range commits {
		v := chart.Value{Value: c.commits, Label: c.date.Format("2006-01-02")}
		values = append(values, v)
	}

	bc := chart.BarChart{
		Width: 1024,
		Title: "Commits by Date",
		Bars:  values,
	}

	buf := bytes.NewBuffer([]byte{})
	err := bc.Render(chart.PNG, buf)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(CommitByDateFileName, buf.Bytes(), 0644)

	return err
}
