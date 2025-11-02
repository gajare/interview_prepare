package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("log.txt")
	defer file.Close()

	name := "Charlie"
	score := 95

	// Arguments are concatenated without spaces: "Name:CharlieScore:95"
	fmt.Fprint(file, "Name:", name, "Score:", score)

	// To add spaces or a newline, you must explicitly include them:
	fmt.Fprint(file, "\n", "New entry.")

	// log.txt content will be:
	// Name:CharlieScore:95
	// New entry.
}
