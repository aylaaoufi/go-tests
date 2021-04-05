package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	table := []string{
		"UD",
		"LL",
	}

	for i := 0; i < 15; i++ {
		table = append(table, rand.RandStr(rand.IntBetween(5, 1000), "UDLR"))
	}

	for _, arg := range table {
		challenge.Program("robottoorigin", arg)
	}
}
