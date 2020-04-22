package models

import (
	"bufio"
	"io"
)

type Job struct {
	RepoUrl  string
	RepoPath string
}

func (j Job) SaveJob(writer io.Writer) (err error) {
	w := bufio.NewWriter(writer)
	_, err = w.WriteString(" world!")
	w.Flush()
	return
}
