package sudoku

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const apiURL string = "https://www.dailysudoku.com/cgi-bin/sudoku/get_board.pl"

type dict map[string]interface{}

func GetTodaySudoku() Sudoku {
	year, month, day := time.Now().Date()
	reqURL := fmt.Sprintf("%s?year=%d&month=%d&day=%d", apiURL, year, month, day)
	log.Println("Making request to", reqURL)
	resp, err := http.Get(reqURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var target dict

	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		log.Fatalln(err)
	}

	return formatToGrid(target["numbers"].(string))
}

func formatToGrid(numbers string) Sudoku {
	newSudoku := make(Sudoku, 9)
	i, j := 0, 0
	for pos, char := range numbers {
		if pos%9 == 0 {
			newSudoku[i] = make(Row, 9)
		}
		if char == '.' {
			newSudoku[i][j] = 0
		} else {
			newSudoku[i][j] = int(char - '0')
		}

		j = (j + 1) % 9

		if pos%9 == 8 {
			i++
		}
	}
	return newSudoku
}
