// Package pos provides functionality for calculating the trading position.
package pos

type Position struct {
	//Ther price at which to buy or sell
	EntryPrice float64
	//How many shares to buy or sell
	Shares int
	//The price at wicch to exit
	TakeProfitPrice float64
	//The price at wicch to stop my loss
	StopLoss float64
	//Expected final profit
	Profit float64
}

// Calculator is an interface that defines a method for calculating a trading position based on the
// gap percentage and opening price.
type Calculator interface {
	Calculate(gapPercent, openingPrice float64) Position
}
