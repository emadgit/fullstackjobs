package sources

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func ScrapeRemoteOK() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// This will hold the entire HTML of the page after rendering
	var html string

	// Navigate and wait for job elements to load
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://remoteok.com"),
		chromedp.Sleep(2*time.Second), // wait for JS-rendered content
		chromedp.OuterHTML("body", &html),
	)
	if err != nil {
		log.Fatalf("Failed to render page: %v", err)
	}

	// Now parse the HTML
	parseJobsFromHTML(html)
	log.Println("RemoteOK scraping completed.")
}
