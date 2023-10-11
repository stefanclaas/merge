package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var mergedLine strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		// Remove trailing newline characters from folded lines
		line = strings.TrimRight(line, "\r\n")

		// Append the line to the merged line
		mergedLine.WriteString(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Print the merged line
	fmt.Println(mergedLine.String())
}

