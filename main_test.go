package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

// TestGreet verifies that the Greet function returns the expected greeting message.
func TestGreet(t *testing.T) {
	expected := "Hello, World!"
	got := Greet()
	if got != expected {
		t.Errorf("Greet() = %q, want %q", got, expected)
	}
}

// TestTimeGreet verifies that TimeGreet returns the correct greeting for different times of day.
func TestTimeGreet(t *testing.T) {
	tests := []struct {
		name     string
		hour     int
		expected string
	}{
		{"midnight", 0, "Good night"},
		{"early morning", 4, "Good night"},
		{"morning start", 5, "Good morning"},
		{"mid morning", 9, "Good morning"},
		{"late morning", 11, "Good morning"},
		{"noon", 12, "Good afternoon"},
		{"mid afternoon", 14, "Good afternoon"},
		{"late afternoon", 16, "Good afternoon"},
		{"evening start", 17, "Good evening"},
		{"mid evening", 19, "Good evening"},
		{"late evening", 20, "Good evening"},
		{"night start", 21, "Good night"},
		{"late night", 23, "Good night"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTime := time.Date(2024, 1, 1, tt.hour, 0, 0, 0, time.UTC)
			got := TimeGreet(testTime)
			if got != tt.expected {
				t.Errorf("TimeGreet(%d:00) = %q, want %q", tt.hour, got, tt.expected)
			}
		})
	}
}

// TestMainOutput verifies that the main function prints the expected output to stdout.
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

	got := buf.String()

	// Check that output contains "Hello, World!"
	if !strings.Contains(got, "Hello, World!") {
		t.Errorf("main() output should contain %q, got %q", "Hello, World!", got)
	}

	// Check that output contains a valid time greeting
	validGreetings := []string{"Good morning", "Good afternoon", "Good evening", "Good night"}
	hasTimeGreeting := false
	for _, greeting := range validGreetings {
		if strings.Contains(got, greeting) {
			hasTimeGreeting = true
			break
		}
	}
	if !hasTimeGreeting {
		t.Errorf("main() output should contain a time greeting, got %q", got)
	}
}
