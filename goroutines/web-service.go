package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// open request stock price for Google
	resp, err := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=googl")
	if err != nil {
		println("Oops, something went wrong with http:", err)
	}

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
