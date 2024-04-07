package main

import (
	"fmt"
	"internal/fetch"
)

const (
	userAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:40.0) Gecko/20100101 Firefox/40.0" +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36"
	pageUrl     = "https://www.otodom.pl/pl/oferty/sprzedaz/mieszkanie/"
	propertyUrl = "https://www.otodom.pl/"
)

func main() {
	results := fetch.FetchListings("warszawa", 3)
	for url, res := range results {
		fmt.Println("URL:", url)
		fmt.Println(res)
		// for _, listing := range res.Listings {
		// 	fmt.Println(propertyUrl+listing, "\n")
		// }
		fmt.Println()
	}
}
