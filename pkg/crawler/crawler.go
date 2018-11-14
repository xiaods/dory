package crawler

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	// "strings"
)

type object struct {
	FinalBalance  uint64 `json:"final_balance"`
	NTX           uint64 `json:"n_tx"`
	TotalReceived uint64 `json:"total_received"`
}

type item struct {
	ChainType    string
	TokenCode    string
	TokenName    string
	ContractAddr string
	TokenIconURL string
	Decimals     int64
}

// TrackingETHContract is token tracking crawler
func TrackingETHContract() error {
	log.Println("Etherscan Contract starting...")

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	c.OnHTML("div#ContentPlaceHolder1_divresult a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://etherscan.io/tokens")

	log.Println("Scraping finished")

	return nil
}

//TrackingETHBalance fetch address balance by interval minutes
func TrackingETHBalance() error {
	log.Println("fetch specified address balance")
	//testing address: 14AcKEr19gHJvgwQhK7sfFm6YJGmoZZoqu
	c := colly.NewCollector(
		//colly.CacheDir("./_instagram_cache/"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	log.Println("starting...")
	c.OnResponse(func(r *colly.Response) {
		jsonData := r.Body

		//parse bitcoin json
		jsonMap := make(map[string]object)
		err := json.Unmarshal([]byte(jsonData), &jsonMap)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("==================")
		for k, v := range jsonMap {
			fmt.Printf("Addr=%+v\n", k)
			fmt.Printf("Balance=%+v\n", float64(v.FinalBalance)/1e8)
		}
		fmt.Println("==================")

	})

	c.Visit("https://blockchain.info/balance?active=14AcKEr19gHJvgwQhK7sfFm6YJGmoZZoqu")

	log.Println("Tracking token balance done....")

	return nil
}
