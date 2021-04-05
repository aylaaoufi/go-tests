package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	// 15 unvalid strings in the table
	table := lib.MultRandASCII()

	// 15 valid strings in the table
	for i := 0; i < 15; i++ {
		table = append(table, strconv.Itoa(lib.RandIntBetween(0, 1000000)))
	}

	// Special cases added to table
	table = append(table,
		"Hello! How are you?",
		"a",
		"z",
		"!",
		"HelloHowareyou",
		"What's this 4?",
		"Whatsthis4",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!",
		"0123456789",
		"010203",
		"01,02,03",
	)
	for _, arg := range table {
		lib.Function("IsNumeric", student.IsNumeric, solutions.IsNumeric, arg)
	}
}
