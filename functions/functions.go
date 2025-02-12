package functions

import (
	"Desktop/chess_assignment/models"
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func GetRowsAndColumns() ([]int, []string) {
	rows := []int{8, 7, 6, 5, 4, 3, 2, 1}
	columns := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	return rows, columns
}
func BuildBoard() []models.Cell {
	board := []models.Cell{}
	rows, columns := GetRowsAndColumns()
	for _, row := range rows {
		for _, column := range columns {
			board = append(board, models.Cell{Row: row, Column: column})
		}
	}
	return board
}

func FindRowColCells(selectedCells []string) ([]models.Cell, error) {
	bishops := []models.Cell{}
	for _, selected := range selectedCells {
		rowCol := strings.Split(selected, "")
		if len(rowCol) == 2 {
			rowConverted, err := strconv.Atoi(rowCol[1])
			if err != nil {
				return nil, errors.New("incorrect input: please follow the example text -> A1 B2 C3 etc")
			}
			bishops = append(bishops, models.Cell{Row: rowConverted, Column: rowCol[0]})
		} else {
			return nil, errors.New("incorrect input: please follow the example text -> A1 B2 C3 etc")
		}
	}
	return bishops, nil
}

func FindSightLine(pieceLocations []models.Cell, board []models.Cell) (bool, []models.Cell, error) {
	for _, pieceLocation := range pieceLocations {
		if !slices.Contains(board, pieceLocation) {
			return false, nil, errors.New("incorrect input: please follow the example text -> A1 B2 C3 etc")
		}
	}

	// get both bishops out of the list to simplify logic and readability
	firstBishopLocation := pieceLocations[0]
	secondBishopLocation := pieceLocations[1]

	// early return example where bishops can never attack
	if firstBishopLocation.Row == secondBishopLocation.Row || firstBishopLocation.Column == secondBishopLocation.Column {
		return false, nil, nil
	}

	// the logic here is first determining that bishops can't be on the same line or column with code above.
	// Then we find the lowest Row piece and go diagonal top left and top right to find possible locations
	// of attack and check these against the location of the other piece.
	var (
		attackLocationFoundLeft   bool
		attackLocationFoundRight  bool
		startingLocation          models.Cell
		attackLocation            models.Cell
		trackPathOfAttackTopLeft  []models.Cell
		trackPathOfAttackTopRight []models.Cell
	)

	if firstBishopLocation.Row < secondBishopLocation.Row {
		startingLocation = firstBishopLocation
		attackLocation = secondBishopLocation
	} else {
		startingLocation = secondBishopLocation
		attackLocation = firstBishopLocation
	}

	rows, columns := GetRowsAndColumns()
	for _, row := range rows {
		// skip creating go routines for rows below or equal to startinglocation row
		if row <= startingLocation.Row {
			continue
		}

		// determine amount of rows away from starting location to find the diagonal
		absoluteDifference := int(math.Abs(float64(row - startingLocation.Row)))
		ColIndex := 0
		for index, column := range columns {
			if column == startingLocation.Column {
				ColIndex = index
			}
		}
		if (ColIndex + absoluteDifference) <= len(columns)-1 {
			topRightAttack := models.Cell{Row: row, Column: columns[ColIndex+absoluteDifference]}
			if attackLocation == topRightAttack {
				attackLocationFoundRight = true
			} else {
				trackPathOfAttackTopRight = append(trackPathOfAttackTopRight, topRightAttack)
			}
		}
		if (ColIndex - absoluteDifference) >= 0 {
			topLeftAttack := models.Cell{Row: row, Column: columns[ColIndex-absoluteDifference]}
			if attackLocation == topLeftAttack {
				attackLocationFoundLeft = true
			} else {
				trackPathOfAttackTopLeft = append(trackPathOfAttackTopLeft, topLeftAttack)
			}
		}
	}

	if attackLocationFoundLeft {
		return attackLocationFoundLeft, trackPathOfAttackTopLeft, nil
	} else if attackLocationFoundRight {
		return attackLocationFoundRight, trackPathOfAttackTopRight, nil
	} else {
		return false, nil, nil
	}
}

func DrawBoard(bishops []models.Cell, highLightTiles []models.Cell, board []models.Cell) {
	stringBoard := ""
	currentRow := 0

	// get both bishops out of the list to simplify logic and readability
	firstBishopLocation := bishops[0]
	secondBishopLocation := bishops[1]
	bishopCell := "🟩"
	attackDiagonal := "🟥"

	// we initialize lastcolor as Black so we can start writing the top left of the board as white and switch from there
	lastColor := "Black"
	alreadySetColor := false
	for i, boardCell := range board {
		if i != 0 && i%8 == 0 {
			stringBoard = stringBoard + "\n"
			// switch colors on the end of the lines so next line starts with the same color
			_ = addCellColor(&lastColor)
			alreadySetColor = false
		}
		if currentRow != boardCell.Row {
			stringBoard = stringBoard + fmt.Sprintf("%v", boardCell.Row)
			currentRow = boardCell.Row
		}
		// check if this should be a bishop tile
		if (firstBishopLocation.Column == boardCell.Column && firstBishopLocation.Row == boardCell.Row) ||
			(secondBishopLocation.Column == boardCell.Column && secondBishopLocation.Row == boardCell.Row) {
			// change color in the background, ignore return value
			if !alreadySetColor {
				_ = addCellColor(&lastColor)
			} else {
				alreadySetColor = false
			}
			stringBoard = stringBoard + bishopCell
		} else if slices.Contains(highLightTiles, boardCell) {
			if !alreadySetColor {
				_ = addCellColor(&lastColor)
			} else {
				alreadySetColor = false
			}
			stringBoard = stringBoard + attackDiagonal
		} else {
			stringBoard = stringBoard + addCellColor(&lastColor)
		}
	}
	fmt.Println(stringBoard)
	// add legend line at the bottom
	fmt.Println("X|A|B|C|D|E|F|G|H|")
}

func addCellColor(lastColor *string) string {
	whiteCell := "⬜"
	blackCell := "⬛"
	if *lastColor == "Black" {
		*lastColor = "White"
		return whiteCell
	} else if *lastColor == "White" {
		*lastColor = "Black"
		return blackCell
	}
	return ""
}
