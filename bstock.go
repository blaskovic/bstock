package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
)

var data = `
description: bu
stocks:
  STOCK:
    ticker: BAASTOCK
    url: url1
    notes: Lol
  UPL:
    ticker: BAASTOCK
    notes: Lol
    url: url2
`

type T struct {
	Stocks map[string]Stock
}

type Stock struct {
	Url      string
	Notes    string
	Currency string
	BuyPrice float64
}

func PriceToString(price float64) string {
	return strconv.FormatFloat(price, 'f', 2, 64)
}

// GetPrice downloads URL and returns it
func GetPrice(url string) float64 {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
		return 0.0
	}

	text := doc.Find("#ctl00_BCPP_Celkem_dvCelkem td.num").First().Text()
	text = strings.Replace(text, ",", ".", -1)
	reg, _ := regexp.Compile("[^0-9.]+")
	text = reg.ReplaceAllString(text, "")

	price, err := strconv.ParseFloat(text, 32)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	return price

}

func main() {
	fmt.Printf("bstock by Branislav Blaskovic\n")

	// Config
	path, _ := filepath.Abs("./stocks.yml")
	yamlFile, errFile := ioutil.ReadFile(path)
	if errFile != nil {
		log.Fatalf("Error: %v", errFile)
	}

	t := T{}
	err := yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//fmt.Printf("%#v\n", t)

	// Table to print
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Ticker", "Price"})

	// Stocks cycler
	for ticker, data := range t.Stocks {
		price := GetPrice(data.Url)
		priceStr := PriceToString(price)
		priceStr = strings.Join([]string{priceStr, data.Currency}, " ")
		table.Append([]string{ticker, priceStr})

	}

	table.SetAlignment(tablewriter.ALIGN_RIGHT)
	table.Render()

	// BAASTOCK
}
