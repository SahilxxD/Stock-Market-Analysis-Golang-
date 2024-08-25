// Package process provides functionality for processing stock data.

package process

import (
	"math"
	"slices"

	"github.com/SahilxxD/Stock-Market-Analysis-Golang-/internal/raw"
)

// filterer is a struct that used to filter raw stocks.
type filterer struct {
	minGap float64
}

// Filter filters the given list of stock candidates based on the minimum gap value.
// It returns the filtered list of stocks.
func (n *filterer) Filter(candidates []raw.Stock) (filtered []raw.Stock) {

	filtered = slices.DeleteFunc(candidates, func(s raw.Stock) bool {
		return math.Abs(s.Gap) < n.minGap
	})

	return
}

// NewFilterer creates a new filterer object with the given minimum gap value.
// It returns a raw.Filterer interface.
func NewFilterer(minGap float64) raw.Filterer {
	return &filterer{
		minGap: minGap,
	}
}
