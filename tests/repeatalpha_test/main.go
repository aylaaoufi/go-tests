package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	args := []string{
		"",
		"Hello",
		"World",
		"Home",
		"Theorem",
		"Choumi is the best cat",
		"abracadaba 01!",
		"abc",
		"MaTheMatiCs",
	}

	args = append(args, rand.MultRandAlnum()...)

	for _, v := range args {
		challenge.Program("repeatalpha", v)
	}
	challenge.Program("repeatalpha")
	challenge.Program("repeatalpha", "", "")
}
