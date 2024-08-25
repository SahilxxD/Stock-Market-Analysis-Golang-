package cmd

import (
	"fmt"
	"log"

	"github.com/SahilxxD/Stock-Market-Analysis-Golang-/internal/news"
	"github.com/SahilxxD/Stock-Market-Analysis-Golang-/internal/pos"
	"github.com/SahilxxD/Stock-Market-Analysis-Golang-/internal/raw"
	"github.com/SahilxxD/Stock-Market-Analysis-Golang-/internal/trade"
)

func Run(ldr raw.Loader, f raw.Filterer, c pos.Calculator, fet news.Fetcher, del trade.Deliverer) error {
	stocks, err := ldr.Load()
	if err != nil {
		log.Println("Error loading stocks", err)
		return fmt.Errorf("error loading stocks: %w", err)
	}

	stocks = f.Filter(stocks)

	selectionChan := make(chan trade.Selection, len(stocks))

	for _, stock := range stocks {
		go func(s raw.Stock, selected chan<- trade.Selection) {

			position := c.Calculate(stock.Gap, s.OpeningPrice)

			articles, err := fet.Fetch(s.Ticker)
			if err != nil {
				log.Printf("Error fetching news for %s: %v\n", stock.Ticker, err)
				selected <- trade.Selection{}
				return
			} else {
				log.Printf("Found %d articles about %s\n", len(articles), stock.Ticker)
			}

			sel := trade.Selection{Ticker: s.Ticker, Position: position, Articles: articles}
			selected <- sel
		}(stock, selectionChan)
	}

	var selections []trade.Selection

	for sel := range selectionChan {
		selections = append(selections, sel)
		if len(selections) == len(stocks) {
			close(selectionChan)
		}
	}

	outputPath := "./opg.json"

	err = del.Deliver(selections)
	if err != nil {
		log.Printf("Error dilevering output to %s: %v\n", outputPath, err)
		return fmt.Errorf("error delivering selections: %w", err)
	}
	return nil
}
