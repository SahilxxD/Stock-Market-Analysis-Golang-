// Package raw provides functionality for working with raw stock data.
package raw

// Stock represents a stock with its ticker symbol, gap percentage, and opening price.
// This might come from a CSV file.

type Stock struct {
	Ticker       string
	Gap          float64
	OpeningPrice float64
}

// Loader is an interface for loading stock data.
type Loader interface {
	Load() ([]Stock, error)
}

// Filterer is an interface for filtering raw stock data.
type Filterer interface {
	Filter([]Stock) []Stock
}
