package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// TestGreet verifies that the Greet function returns the expected greeting message.
func TestGreet(t *testing.T) {
	expected := "Hello, World!"
	got := Greet()
	if got != expected {
		t.Errorf("Greet() = %q, want %q", got, expected)
	}
}

// TestMainOutput verifies that the main function prints "Hello, World!" to stdout.
func TestMainOutput(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w

	// Run main
	main()

	// Restore stdout and read captured output
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Fatalf("failed to read captured output: %v", err)
	}

	expected := "Hello, World!\n"
	got := buf.String()
	if got != expected {
		t.Errorf("main() output = %q, want %q", got, expected)
	}
}
