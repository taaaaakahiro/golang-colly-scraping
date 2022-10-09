package main

import (
	"github.com/taaaaakahiro/golang-colly-scraping/perse"
)

var asinCodes = []string{
	"B07CZHKDQ6",
	"B07TCH4JFR",
}

func main() {
	run()
}

func run() {
	// init
	p := perse.NewPerse()

	// func
	// p.Example.Scraping()
	p.Amazon.Scraping(asinCodes)

}
