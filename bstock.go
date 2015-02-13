// bstock
// Author: Branislav Blaskovic

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

type T struct {
	Stocks map[string]Stock
}

type Stock struct {
	Url      string
	Notes    string
	Currency string
	BuyPrice string
	Amount   int
	Fees     float64
}

func FloatToString(number float64) string {
	return strconv.FormatFloat(number, 'f', 2, 64)
}

func StringToFloat(text string) float64 {
	text = strings.Replace(text, ",", ".", -1)
	reg, _ := regexp.Compile("[^0-9.]+")
	text = reg.ReplaceAllString(text, "")

	num, err := strconv.ParseFloat(text, 32)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	return num
}

func JoinStrings(str1, str2 string) string {
	return strings.Join([]string{str1, str2}, " ")
}

func PercentDifference(num1, num2 float64) float64 {
	return (num1 - num2) / ((num1 + num2) / 2)

}

// GetPrice downloads URL and returns it
func GetPrice(url string) float64 {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
		return 0.0
	}

	var text string

	switch {
	case strings.Contains(url, "bcpp.cz") == true:
		text = doc.Find("#ctl00_BCPP_KontinualOL_dvTable td.num").First().Text()
	case strings.Contains(url, "google.com/finance"):
		text = doc.Find("span.pr span").First().Text()
	}

	return StringToFloat(text)

}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Error: argument has to be yaml file with stocks")
	}

	// Config
	path, _ := filepath.Abs(os.Args[1])
	yamlFile, errFile := ioutil.ReadFile(path)
	if errFile != nil {
		log.Fatalf("Error: %v", errFile)
	}

	t := T{}
	err := yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// This is just for debug
	// fmt.Printf("%#v\n", t)

	// Table to print
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Ticker", "Price", "Buy price", "Diff", "Diff %", "Fees", "Overall", "Notes"})

	// Stocks cycler
	for ticker, data := range t.Stocks {
		// Prepare variables to display
		dTicker := JoinStrings(strconv.Itoa(data.Amount), ticker)
		price := GetPrice(data.Url)
		buyPrice := StringToFloat(data.BuyPrice)
		dPrice := JoinStrings(FloatToString(price), data.Currency)
		dBuyPrice := JoinStrings(data.BuyPrice, data.Currency)
		dDifference := JoinStrings(FloatToString(price-buyPrice), data.Currency)
		dDifferencePercent := JoinStrings(FloatToString(PercentDifference(price, buyPrice)), "%")
		dOverall := JoinStrings(FloatToString(float64(data.Amount)*(price-buyPrice)-data.Fees), data.Currency)
		dFees := JoinStrings(FloatToString(data.Fees), data.Currency)

		// Append row to table
		table.Append([]string{dTicker, dPrice, dBuyPrice, dDifference, dDifferencePercent, dFees, dOverall, data.Notes})
	}

	table.SetAlignment(tablewriter.ALIGN_RIGHT)
	table.Render()
}
