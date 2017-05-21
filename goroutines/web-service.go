package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

func main() {

	// make 4 cores available to this app
	runtime.GOMAXPROCS(4)

	// keep track of how long web service requests take
	start := time.Now()

	// list of stock symbols
	stockSymbols := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}

	// tracking variable
	numComplete := 0

	// lookup price for each company
	for _, symbol := range stockSymbols {
		// concurrent goroutine, notice it gets its own version of symbol
		// to decouple it from main loop execution
		go func(symbol string) {
			// open request stock price for Google
			resp, _ := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=" + symbol)

			// close connection when we're done
			defer resp.Body.Close()

			// get content of entire response into a single byte slice
			body, _ := ioutil.ReadAll(resp.Body)

			// create QuoteResponse object to hold result of web service call
			quote := new(QuoteResponse)

			// convert raw http body response to QuoteResponse type
			xml.Unmarshal(body, &quote)

			// print name of company and stock price to console
			fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
			numComplete++
		}(symbol) // invoke goroutine with symbol variable from main loop
	}

	// wait for a bit if there's still work to do
	for numComplete < len(stockSymbols) {
		time.Sleep(10 * time.Millisecond)
	}

	// calculate how long requests took
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)

}

// struct definition that mirrors the xml doc
type QuoteResponse struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	Timestamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChangeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}
