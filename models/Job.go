package models

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

type Job struct {
	RepoUrl  string
	RepoPath string
}

func (j Job) WriteJob(writer io.Writer) (err error) {
	data, _ := json.MarshalIndent(j, "", " ")
	w := bufio.NewWriter(writer)
	_, err = w.Write(data)
	w.Flush()
	return
}

func (j Job) PersistJob() (err error) {
	file, err := os.Create("./test.txt")
	MyJob := Job{"repo", "path/in/repo"}
	err = MyJob.WriteJob(file)
	return
}
