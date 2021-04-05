package main

import (
	"strings"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{"Hello how are you?"}

	// 30 random slice of strings
	for i := 0; i < 30; i++ {
		table = append(table, strings.Join(rand.MultRandASCII(), " "))
	}
	for _, arg := range table {
		challenge.Function("SplitWhiteSpaces", student.SplitWhiteSpaces, solutions.SplitWhiteSpaces, arg)
	}
}
