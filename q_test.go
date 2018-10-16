package q

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWord(t *testing.T) {
	assert.Equal(t, Word("word1 word2 word3 word4\t\n"), []string{"word1", "word2", "word3", "word4"})
	assert.Equal(t, Word("  word1 word2 word3 word4  \t\n"), []string{"word1", "word2", "word3", "word4"})
}

func TestLine(t *testing.T) {
	assert.Equal(t, Line(`
    # Comment is ignored
    // comment is ignored

    line 1 
    line 2
    line 3

  `), []string{"line 1", "line 2", "line 3"})
}

func TestHereDoc(t *testing.T) {
	assert.Equal(t, HereDoc(`
    First Line no indent
      Second line indent is kept
        third line indent 4
  `), "First Line no indent\n  Second line indent is kept\n    third line indent 4\n")
}

func TestCmd(t *testing.T) {
	assert.Equal(t, Cmd(`
    // construct command lines. Same as Line(), join lines with " && " or other connector like ";" 
    ls -ltr

    hostname
  `, " && "), "ls -ltr && hostname")
}

func TestTemplate(t *testing.T) {
	text := Template(
		`This is a test for {{ .Name}} at {{ .Address }}`,
		map[string]string{
			"Name":    "test",
			"Address": "Street 1",
		})

	assert.Equal(t, text, "This is a test for test at Street 1")
}
