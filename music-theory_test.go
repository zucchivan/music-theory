// Package main implements a command-line utility for music
package main

import (
	"os"
	"testing"
)

func TestMusicTheory(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd",
		"chord", "G minor 7",
	}
	main()
}

func TestHelpCmd(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd",
		"-h",
	}
	main()
}
