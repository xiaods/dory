package crawler

import (
	"github.com/gocolly/colly"
	"log"
)

type item struct {
	ChainType    string
	TokenCode    string
	TokenName    string
	ContractAddr string
	TokenIconURL string
	Decimals     int64
}

//TrackingETHContract is token tracking crawler
func TrackingETHContract() error {
	log.Println("Etherscan Contract starting...")

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	knownUrls := []string{}
	c.OnHTML("html", func(e *colly.HTMLElement) {
		log.Println(e)
		log.Println("============")
		knownUrls = append(knownUrls, e.ChildText("a"))
	})

	c.Visit("https://etherscan.io/tokens?p=1")

	log.Println(knownUrls)
	log.Println("Scraping finished")

	return nil
}
