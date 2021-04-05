package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib"
)

func main() {
	table := append(lib.MultRandWords(), "1 2 3 4 5")

	for _, s := range table {
		lib.Program("boolean", strings.Fields(s)...)
	}
}
