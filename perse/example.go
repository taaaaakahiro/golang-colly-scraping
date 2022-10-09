package perse

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Example struct {
}

func NewExample() *Example {
	return &Example{}
}

func (s *Example) Scraping() {
	p := &pageInfo{}
	articles := make([]*item, 0, 4)

	// Instantiate default collector
	c := colly.NewCollector()

	// Extract title element
	c.OnHTML("title", func(e *colly.HTMLElement) {
		p.Title = e.Text
		fmt.Println(e.Text)
	})

	i := 0
	// Extract li class="new-entry-item"
	c.OnHTML("li[class=new-entry-item]", func(e *colly.HTMLElement) {
		i++
		fmt.Println(i)

		// Extract h3 element
		title := e.ChildText("h3")
		fmt.Println(title)

		// Extract href
		link, _ := e.DOM.Find("a[href]").Attr("href")
		fmt.Println(link)

		article := &item{
			Title: title,
			URL:   link,
		}

		articles = append(articles, article)
		p.Article = articles
	})

	// Before making a request print "Visiting URL: https://XXX"
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

	// Start scraping on https://XXX
	c.Visit(ExUrl)

	// Wait until threads are finished
	c.Wait()

	// Save as JSON format
	fileName := fmt.Sprintf("./%s/%s.json", dir, fileName)
	savePageJson(fileName, p)
}
