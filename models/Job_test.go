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
