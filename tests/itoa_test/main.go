package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	for i := 0; i < 50; i++ {
		lib.Function("Itoa", student.Itoa, solutions.Itoa, lib.RandInt())
	}
}
