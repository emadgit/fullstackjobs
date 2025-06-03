package sources

import (
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func ScrapeRemoteOK() {
	c := colly.NewCollector(
		colly.AllowedDomains("remoteok.com"),
	)

	c.OnHTML("tr.job", func(e *colly.HTMLElement) {
		title := e.ChildText("h2")
		company := e.ChildText("h3")
		location := e.ChildText("span.location")
		datePosted := e.ChildText("time")

		if title == "" || company == "" {
			return // Skip if title or company is missing
		}

		log.Printf("Job Title: %s, Company: %s, Location: %s, Date Posted: %s\n",
			strings.TrimSpace(title), strings.TrimSpace(company), strings.TrimSpace(location), strings.TrimSpace(datePosted))
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request failed with status code %d: %v", r.StatusCode, err)
	})

	err := c.Visit("https://remoteok.com/")
	if err != nil {
		log.Fatalf("Failed to visit RemoteOK: %v", err)
	}

	time.Sleep(2 * time.Second) // Allow time for scraping to complete
	log.Println("RemoteOK scraping completed.")
}
