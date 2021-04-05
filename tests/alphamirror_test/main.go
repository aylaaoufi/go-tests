package main

import (
	"github.com/01-edu/go-tests/lib"
)

func main() {
	args := [][]string{
		{""},
		{"One", "ring!"},
		{"testing spaces and #!*"},
		{"more", "than", "three", "arguments"},
		{"Upper anD LoWer cAsE"},
		{lib.RandWords()},
	}

	for _, v := range args {
		lib.Program("alphamirror", v...)
	}
}
