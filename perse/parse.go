package perse

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/oklog/ulid/v2"
)

var (
	fileName = ulid.Make()
)

const (
	dir   string = "output"
	ExUrl string = "https://cpp-learning.com"
	Url          = "https://www.amazon.co.jp/dp/"
)

type item struct {
	Title string `json:"titleddd"`
	URL   string `json:"url"`
	Price int    `json"price"`
}

type pageInfo struct {
	StatusCode int     `json:"statusCode"`
	URL        string  `json:"url"`
	Title      string  `json:"title"`
	Article    []*item `json:"article"`
	Price      int     `json:"price"`
}

type Parse struct {
	Example *Example
	Amazon  *Amazon
}

func NewPerse() *Parse {
	return &Parse{
		Example: NewExample(),
	}
}

func savePageJson(fName string, p *pageInfo) {
	// Create json file
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	// Dump json to the standard output
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	err = enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}

	// Struct to json
	b, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println(string(b))
}
