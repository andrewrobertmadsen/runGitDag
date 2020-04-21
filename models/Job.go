package models

import "io"

type Job struct {
	RepoUrl   string
	RepoPath   string
}

func (j Job) SaveJob(writer io.Writer) (err error) {
	writer.Write("hello")

	return
}
