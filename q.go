package q

import (
	"regexp"
	"strings"
)

// Word return the string slice of the words
func Word(line string) []string {
	return regexp.MustCompile(`\s+`).Split(line, -1)
}

// Line return the lines as a string slice, ignoring empty line, line as comment (started with // or #)
func Line(lines string) []string {
	var result []string
	for _, l := range regexp.MustCompile(`\n`).Split(lines, -1) {
		nl := strings.TrimSpace(l)
		if nl == "" || strings.HasPrefix(nl, "//") || strings.HasPrefix(nl, "#") {
			continue
		}
		result = append(result, nl)
	}

	return result
}
