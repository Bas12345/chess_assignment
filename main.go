package main

import (
	"Desktop/chess_assignment/functions"
	"fmt"
	"log"
)

func main() {
	var firstBishop, secondBishop string
	fmt.Println("This program determines if two bishops on a chess board can attack eachother. Please enter the chosen location of your bishops by row and column. Input example: A1 or B4")
	fmt.Println("Enter first bishop location: ")
	fmt.Scanln(&firstBishop)
	fmt.Println("Enter second bishop location: ")
	fmt.Scanln(&secondBishop)
	selectedCells, err := functions.FindRowColCells([]string{firstBishop, secondBishop})
	if err != nil {
		log.Fatal(err)
	}
	board := functions.BuildBoard()
	found, highLightTiles, err := functions.FindSightLine(selectedCells, board)
	if err != nil {
		log.Fatal(err)
	}
	if found {
		fmt.Println("Bishops can attack eachother.")
		fmt.Println(highLightTiles)
		functions.DrawBoard(selectedCells, highLightTiles, board)
	} else {
		fmt.Println("Bishops can not attack eachother.")
		functions.DrawBoard(selectedCells, highLightTiles, board)
	}
}
