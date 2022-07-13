package sudoku

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SolveOnWebsite(sudoku Sudoku) {
	// open chrome in headless
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	err := chromedp.Run(ctx, chromedp.Tasks{
		// takes long time to load, and is inconsequential for the rest of the page elements
		network.SetBlockedURLS([]string{"https://www.assoc-amazon.co.uk/e/ir?t=dailysudoku-21&l=as2&o=2&a=1897597649"}),
		chromedp.Navigate("https://www.dailysudoku.com/sudoku/play.shtml?today=1"),
		chromedp.WaitVisible("form", chromedp.ByQuery),
	})

	checkErr(err)

	for i, row := range sudoku {
		for j, cell := range row {
			cell_num := i*9 + j
			log.Println("Putting", cell, "to", cell_num)
			err = chromedp.Run(ctx, chromedp.Tasks{
				chromedp.Click(fmt.Sprint("#p", cell_num), chromedp.ByID),
				chromedp.SendKeys(fmt.Sprint("#c", cell_num), strconv.Itoa(cell), chromedp.ByID),
			})
			checkErr(err)
		}
	}
}
