package sources

import (
	"context"
	"log"
	"time"

	"fullstackjobs-service/cmd/internal/storage"

	"github.com/chromedp/chromedp"
)

func ScrapeRemoteOK() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var html string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://remoteok.com"),
		chromedp.Sleep(2*time.Second),
		chromedp.OuterHTML("body", &html),
	)
	if err != nil {
		log.Fatalf("Failed to render page: %v", err)
	}

	jobs := parseJobsFromHTML(html)

	log.Printf("✅ Parsed %d jobs from RemoteOK", len(jobs))

	if len(jobs) > 0 {
		err := storage.SaveJobs(jobs)
		if err != nil {
			log.Printf("❌ Failed to save jobs: %v", err)
		} else {
			log.Println("✅ Jobs saved to DB.")
		}
	} else {
		log.Println("⚠️ No new jobs to save.")
	}

	log.Println("RemoteOK scraping completed.")
}

