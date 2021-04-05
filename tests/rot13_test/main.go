package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	table := append(rand.MultRandWords(),
		" ",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ",
		"a b c d e f g h ijklmnopqrstuvwxyz A B C D E FGHIJKLMNOPRSTUVWXYZ",
	)
	for _, s := range table {
		challenge.Program("rot13", s)
	}
	challenge.Program("rot13", "1 argument", "2 arguments")
	challenge.Program("rot13", "1 argument", "2 arguments", "3 arguments")
}
