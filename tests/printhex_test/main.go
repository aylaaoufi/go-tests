package main

import (
	"strconv"
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	table := []string{
		" ",
		"123 132 1",
		"1 5",
		"0",
	}
	for i := 0; i < 10; i++ {
		table = append(table, strconv.Itoa(rand.IntBetween(-1000, rand.MaxInt)))
	}
	for i := 0; i < 15; i++ {
		table = append(table, strconv.Itoa(i))
	}
	for _, s := range table {
		challenge.Program("printhex", strings.Fields(s)...)
	}
}
