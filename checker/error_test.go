package checker_test

import (
	"testing"

	"github.com/vodkabears/yaspell/checker"
	"github.com/vodkabears/yaspell/reader"
)

func TestCheckerError(t *testing.T) {
	misspells := []checker.Misspell{
		{
			Code: 1,
			Pos:  0,
			Row:  0,
			Col:  0,
			Len:  6,
			Word: "Golang",
		},
		{
			Code:        1,
			Pos:         0,
			Row:         1,
			Col:         0,
			Len:         6,
			Word:        "Goroutines",
			Suggestions: []string{"Go routines"},
		},
	}
	err := checker.Error{
		&reader.Chunk{"file.txt", "test"},
		misspells,
	}
	expected := "file.txt:1:1: Golang\n" +
		"file.txt:2:1: Goroutines [Go routines]"

	if err.Error() != expected {
		t.Errorf("\nExpected to get:\n%s.\nGot:\n%s", err, expected)
	}
}
