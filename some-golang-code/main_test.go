package main

import (
	"testing"
	"os"
	"bytes"
	"io"
)

func TestMain(t *testing.T) {
	// capture stdout
  old_stdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Error creating pipe: %s", err)
	}
	os.Stdout = w
	defer func() {
		os.Stdout = old_stdout
	}()
	outChannel := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outChannel <- buf.String()
	}()
	main()
	w.Close()
	// compare stdout with expected output
	out := <-outChannel
	if out != "Hello, World!!\n" {
		t.Errorf("Expected 'Hello, World!!', got '%s'", out)
	}
}