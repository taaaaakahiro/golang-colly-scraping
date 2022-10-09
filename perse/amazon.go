package perse

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Amazon struct{}

func NewAmazon() *Amazon {
	return &Amazon{}
}

func (a *Amazon) Scraping(asinCodes []string) {
	p := &pageInfo{}

	// Instantiate default collector
	c := colly.NewCollector()

	// Extract title element
	c.OnHTML("title", func(e *colly.HTMLElement) {
		p.Title = e.Text
		fmt.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		p.URL = r.URL.String()
		fmt.Println("Visiting URL:", r.URL.String())
	})

	// After making a request extract status code
	c.OnResponse(func(r *colly.Response) {
		p.StatusCode = r.StatusCode
		fmt.Println("StatusCode:", r.StatusCode)
	})
	c.OnError(func(r *colly.Response, err error) {
		p.StatusCode = r.StatusCode
		log.Println("error:", r.StatusCode, err)
	})

	c.OnHTML("#corePrice_feature_div", func(e *colly.HTMLElement) {
		// price, err := strconv.Atoi(e.Text)
		str := e.ChildText(".a-price-whole")
		fmt.Println(str)
		str = strings.Replace(str, ",", "", -1)

		price, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
			return
		}

		p.Price = price

	})

	// Start scraping on https://XXX
	for _, asinCode := range asinCodes {
		c.Visit(Url + asinCode)
		// Wait until threads are finished
		c.Wait()

		// Save as JSON format
		fileName := fmt.Sprintf("./%s/%s.json", dir, fileName)
		savePageJson(fileName, p)

	}

}
