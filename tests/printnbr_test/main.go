package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(rand.Ints(),
		rand.MinInt,
		rand.MaxInt,
		0,
	)
	for _, arg := range table {
		challenge.Function("PrintNbr", student.PrintNbr, solutions.PrintNbr, arg)
	}
}
