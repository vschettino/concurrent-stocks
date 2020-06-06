package main

import (
	"fmt"
	"sync"
	"time"
	"github.com/markcheno/go-quote"

)

func main() {
	tickers := []string{"FB", "AAPL", "AMZN", "MSFT", "GOOG"}
	from := time.Now().AddDate(-10, 0, 0)
	to := time.Now()
	fetchQuotes(tickers, from, to)
}

func fetchQuotes(tickers []string, from time.Time, to time.Time) error {
	var group sync.WaitGroup
	group.Add(len(tickers))
	for _, ticker := range tickers {
		go fetchQuote(ticker, from, to, &group)
	}
	group.Wait()
	return nil
}

func fetchQuote(ticker string, from time.Time, to time.Time, group *sync.WaitGroup){
	defer group.Done()
	quotes, err := quote.NewQuoteFromYahoo(ticker,
			from.Format("2006-01-02"),
			to.Format("2006-01-02"),
			quote.Daily,
			true,
		)
	totalDataPoints := len(quotes.Close)
	if err != nil {
		fmt.Printf("[%s] Error fetching: %s\n", ticker, err)
		return
	}
	fmt.Printf("[%s] Fetched %d data points\n", ticker, totalDataPoints)
}