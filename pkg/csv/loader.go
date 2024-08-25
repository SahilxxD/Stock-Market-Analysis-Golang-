// Package csv provides functionality for loading and processing CSV files containing stock data.

package csv

import (
	"encoding/csv"
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/SahilxxD/Stock-Market-Analysis-Golang-/internal/raw"
)

// columns represents a slice of strings, where each string represents a column in the CSV file.
type columns = []string

// rows represents a slice of columns, where each columns represents a row in the CSV file.
type rows = []columns

// loader is responsible for loading stock data from a CSV file.
type loader struct {
	path string
}

func (l *loader) read() (rows, error) {
	f, err := os.Open(l.path)

	if err != nil {
		log.Println("Error opening", err)
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		log.Println("Error reading csv", err)
		return nil, err
	}
	rows = slices.Delete(rows, 0, 1)
	return rows, nil
}

// Load reads the CSV file and returns a slice of raw.Stock objects representing the stock data.
// It skips rows with invalid data and returns an error if there was a problem reading the file.
func (l *loader) Load() ([]raw.Stock, error) {
	rows, err := l.read()
	if err != nil {
		return nil, err
	}

	var stocks []raw.Stock

	for _, row := range rows {
		ticker := row[0]
		openingPrice, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			log.Printf("Error parsing openingPrice %s: %v\n", row[2], err)
			continue
		}
		gap, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			log.Printf("Error parsing gap %s: %v\n", row[1], err)
			continue
		}

		stocks = append(stocks, raw.Stock{
			Ticker:       ticker,
			Gap:          gap,
			OpeningPrice: openingPrice,
		})
	}
	return stocks, nil
}

// NewLoader creates a new instance of the loader with the specified file path.
func NewLoader(path string) raw.Loader {
	return &loader{
		path: path,
	}
}
