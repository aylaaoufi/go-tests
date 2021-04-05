package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/is"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	functions := []func(string) bool{is.Digit, is.Lower, is.Upper}

	type node struct {
		f func(string) bool
		a []string
	}

	table := []node{}

	for i := 0; i < 5; i++ {
		function := functions[lib.RandIntBetween(0, len(functions)-1)]
		table = append(table, node{
			f: function,
			a: lib.MultRandWords(),
		})
	}
	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: is.Digit,
			a: lib.MultRandDigit(),
		})
	}

	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: is.Lower,
			a: lib.MultRandLower(),
		})
	}
	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: is.Upper,
			a: lib.MultRandUpper(),
		})
	}

	table = append(table,
		node{
			f: is.Digit,
			a: []string{"Hello", "how", "are", "you"},
		},
		node{
			f: is.Digit,
			a: []string{"This", "is", "4", "you"},
		},
	)

	for _, arg := range table {
		lib.Function("Any", student.Any, solutions.Any, arg.f, arg.a)
	}
}
