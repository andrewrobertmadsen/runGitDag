package models

import (
	"bytes"
	"testing"
)

func TestSaveJob(t *testing.T) {
	testJob := Job{"this is a repo url", "path/in/repo"}
	buf := bytes.NewBufferString("Hello")
	err := testJob.SaveJob(buf)
	if err != nil {
		t.Fail()
	}
	got := buf.String()
	want := "Hello world!"

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
