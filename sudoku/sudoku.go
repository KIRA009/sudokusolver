package sudoku

import (
	"fmt"
)

type Row []int

type Sudoku []Row

const NUMCELLS int = 9

func (s Sudoku) hasEmptyCell() bool {
	for i := 0; i < NUMCELLS; i++ {
		for j := 0; j < NUMCELLS; j++ {
			if s[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func (s Sudoku) copy() Sudoku {
	newSudoku := make(Sudoku, len(s))
	for i := range s {
		newSudoku[i] = make(Row, len(s[i]))
		copy(newSudoku[i], s[i])
	}
	return newSudoku
}

func (s Sudoku) PrintBoard() {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < NUMCELLS; row++ {
		fmt.Print("| ")
		for col := 0; col < NUMCELLS; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", s[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func (s Sudoku) Solve(channel *chan Sudoku) bool {
	if !s.hasEmptyCell() {
		*channel <- s
		return true
	}
	for i := 0; i < NUMCELLS; i++ {
		for j := 0; j < NUMCELLS; j++ {
			if s[i][j] == 0 {
				for candidate := NUMCELLS; candidate >= 1; candidate-- {
					s[i][j] = candidate
					if s.isBoardValid() {
						go func(s Sudoku, channel *chan Sudoku, i, j int) {
							s.Solve(channel)
						}(s.copy(), channel, i, j)
						s[i][j] = 0
					} else {
						s[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func (s Sudoku) isBoardValid() bool {

	//check duplicates by row
	for row := 0; row < NUMCELLS; row++ {
		counter := [NUMCELLS + 1]int{}
		for col := 0; col < NUMCELLS; col++ {
			counter[s[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < NUMCELLS; row++ {
		counter := [NUMCELLS + 1]int{}
		for col := 0; col < NUMCELLS; col++ {
			counter[s[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	// check 3x3 section
	for i := 0; i < NUMCELLS; i += 3 {
		for j := 0; j < NUMCELLS; j += 3 {
			counter := [NUMCELLS + 1]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[s[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter [NUMCELLS + 1]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}
