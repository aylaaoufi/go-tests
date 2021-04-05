package main

import (
	"github.com/01-edu/go-tests/lib"
)

func main() {
	table := append(lib.MultRandWords(),
		"",
		"             a as",
		"   f     d",
		"   asd    ad",
		"   salut !!! ",
		" salut ! ! !",
		"salut ! !",
	)
	for _, s := range table {
		lib.Program("firstword", s)
	}
}
