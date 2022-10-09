package main

import (
	"github.com/taaaaakahiro/golang-colly-scraping/perse"
)

func main() {
	run()
}

func run() {
	// init
	p := perse.NewPerse()

	// func
	// p.Example.Scraping()
	p.Amazon.Scraping()

}
