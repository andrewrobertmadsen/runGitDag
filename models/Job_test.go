package models

import (
	"bytes"
	"io"
	"testing"
)

func TestSaveJob(t *testing.T) {
	test_job := Job{"this is a repo url", "path/in/repo"}
	var buf bytes.Buffer
	writer := io.Writer(&buf)
	test_job.SaveJob(writer)
	got := buf.String()
	want := "Hello"

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}

}