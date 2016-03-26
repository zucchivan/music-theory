// Package main implements a command-line utility for music
package main

import (
	"os"
	"testing"
)

func TestSpec(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd",
		"-k", "G minor 7",
	}
	main()
}
