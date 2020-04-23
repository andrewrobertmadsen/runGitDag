package models

import (
	"bytes"
	"testing"
)

func TestWriteJob(t *testing.T) {
	testJob := Job{"url", "path"}
	buf := bytes.NewBufferString("")
	err := testJob.WriteJob(buf)
	if err != nil {
		t.Fail()
	}
	got := buf.String()
	want := "{\n \"RepoUrl\": \"url\",\n \"RepoPath\": \"path\"\n}"

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func TestParseJob(t *testing.T) {
	x := []byte("{\"RepoUrl\": \"url\",\"RepoPath\": \"path\"}")
	got, err := ParseJob(x)
	want := Job{"url", "path"}
	if err != nil {
		t.Fail()
	}
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
