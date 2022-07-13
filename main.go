package main

import (
	"log"
	"time"

	"github.com/KIRA009/sudokusolver/sudoku"
)

func main() {
	s := sudoku.GetTodaySudoku()

	solved := make(chan sudoku.Sudoku)

	now := time.Now()

	go s.Solve(&solved)

	s = <-solved

	log.Println("Time to solve:", time.Since(now))

	s.PrintBoard()

	sudoku.SolveOnWebsite(s)
}
