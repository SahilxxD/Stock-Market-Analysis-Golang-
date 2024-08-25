package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SahilxxD/Stock-Market-Analysis-Golang-/cmd"
	"github.com/SahilxxD/Stock-Market-Analysis-Golang-/internal/news"
	"github.com/SahilxxD/Stock-Market-Analysis-Golang-/internal/pos"
	"github.com/SahilxxD/Stock-Market-Analysis-Golang-/internal/raw"
	"github.com/SahilxxD/Stock-Market-Analysis-Golang-/internal/trade"
	"github.com/SahilxxD/Stock-Market-Analysis-Golang-/pkg/csv"
	"github.com/SahilxxD/Stock-Market-Analysis-Golang-/pkg/json"
	"github.com/SahilxxD/Stock-Market-Analysis-Golang-/pkg/process"
	"github.com/SahilxxD/Stock-Market-Analysis-Golang-/pkg/salpha"
)

func main() {

	var seekingAlphaURL = "https://seeking-alpha.p.rapidapi.com"
	var seekingAlphaAPIKey = "API key"

	// Validate environment variables
	if seekingAlphaURL == "" {
		fmt.Println("Missing SEEKING_ALPHA_URL environment variable")
		os.Exit(1)
	}

	if seekingAlphaAPIKey == "" {
		fmt.Println("Missing SEEKING_ALPHA_API_KEY environment variable")
		os.Exit(1)
	}

	// Define command-line flags
	inputPath := flag.String("i", "", "path to input file (required)")
	accountBalance := flag.Float64("b", 0.0, "Account balance (required)")
	outputPath := flag.String("o", "./opg.json", "Path to output file.")
	lossTolerance := flag.Float64("l", 0.02, "Loss tolerance percentage")
	profitPercent := flag.Float64("p", 0.8, "Percentage of the gap to take as profit")
	minGap := flag.Float64("m", 0.1, "Minimum gap value to consider")

	// Parse command-line flags
	flag.Parse()

	// Check if required flags are provided
	if *inputPath == "" || *accountBalance == 0.0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var ldr raw.Loader = csv.NewLoader(*inputPath)
	var f raw.Filterer = process.NewFilterer(*minGap)
	var c pos.Calculator = process.NewCalculator(*accountBalance, *lossTolerance, *profitPercent)
	var fet news.Fetcher = salpha.NewClient(seekingAlphaURL, seekingAlphaAPIKey)
	var del trade.Deliverer = json.NewDeliverer(*outputPath)

	err := cmd.Run(ldr, f, c, fet, del)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
