package functions

import (
	"Desktop/chess_assignment/models"
	"testing"
)

func Test_FindSightLine(t *testing.T) {
	board := BuildBoard()
	found, path, err := FindSightLine([]models.Cell{{Row: 1, Column: "A"}, {Row: 8, Column: "H"}}, board)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if !found {
		t.Fatalf("possible attack should be found")
	}
	if len(path) == 0 {
		t.Fatalf("attack path should be found")
	}
}

func Test_DoesntFindSightLine(t *testing.T) {
	board := BuildBoard()
	found, path, err := FindSightLine([]models.Cell{{Row: 1, Column: "A"}, {Row: 8, Column: "G"}}, board)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if found {
		t.Fatalf("possible attack should not be found")
	}
	if len(path) > 0 {
		t.Fatalf("attack path should not be found")
	}
}

func Test_IncorrectInputFindSightLine(t *testing.T) {
	board := BuildBoard()
	_, _, err := FindSightLine([]models.Cell{{Row: 9, Column: "Z"}, {Row: 8, Column: "G"}}, board)
	if err == nil {
		t.Fatalf("Should error on input out of bounds of chess board")
	}
}
