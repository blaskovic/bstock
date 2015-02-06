package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"strconv"
	"strings"
)

// Download downloads URL and returns it
func GetPrice(url string) float64 {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	text := doc.Find("#ctl00_BCPP_Celkem_dvCelkem td.num").First().Text()
	text = strings.Replace(text, ",", ".", -1)

	price, err := strconv.ParseFloat(text, 32)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	return price

}

func main() {
	fmt.Printf("bstock by Branislav Blaskovic\n")

	// BAASTOCK
	price := GetPrice("http://www.bcpp.cz/Cenne-Papiry/Detail.aspx?isin=GB00BF5SDZ96")
	fmt.Printf("BAASTOCK %.2f CZK\n", price)
}
