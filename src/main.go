package main

// maybe make this a web server and have a front end?

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("Hello world!")

	urls := []string{"sportsbook.draftkings.com",
		"www.geeksforgeeks.org",
	}

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(urls...),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("", func(e *colly.HTMLElement) {
		// link := e.Attr("class")
		// Print link
		// fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		fmt.Println("Found item %s", e.Text)

		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		// c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on ...
	c.Visit("https://sportsbook.draftkings.com/leagues/basketball/nba")

}
