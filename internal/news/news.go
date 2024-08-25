// Package news provides functionality for fetching stock news.
package news

import "time"

// Article represents a news article.
type Article struct {
	PublishOn time.Time // PublishOn represents the publication date and time of the article.
	Headline  string    // Headline represents the title or headline of the article.
}

// Fetcher is an interface for fetching news articles.
type Fetcher interface {
	// Fetch retrieves news articles for the given ticker symbol.
	// It returns a slice of Article and an error if any.
	Fetch(ticker string) ([]Article, error)
}
