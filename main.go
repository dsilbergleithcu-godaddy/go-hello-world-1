package main

import (
	"fmt"
	"time"
)

// Greet returns the greeting message.
func Greet() string {
	return "Hello, World!"
}

// TimeGreet returns a time-of-day greeting based on the given time.
func TimeGreet(now time.Time) string {
	hour := now.Hour()
	switch {
	case hour >= 5 && hour < 12:
		return "Good morning"
	case hour >= 12 && hour < 17:
		return "Good afternoon"
	case hour >= 17 && hour < 21:
		return "Good evening"
	default:
		return "Good night"
	}
}

func main() {
	fmt.Println(Greet())
	fmt.Println(TimeGreet(time.Now()))
}
