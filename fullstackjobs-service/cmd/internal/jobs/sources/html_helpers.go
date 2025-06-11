package sources

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"fullstackjobs-service/cmd/internal/models"

	"golang.org/x/net/html"
)

func parseJobsFromHTML(htmlStr string) []models.Job {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	var jobs []models.Job

	var RecursiveFn func(*html.Node)
	RecursiveFn = func(n *html.Node) {
		// 💡 Check for <script type="application/ld+json">
		if n.Type == html.ElementNode && n.Data == "script" {
			for _, attr := range n.Attr {
				if attr.Key == "type" && attr.Val == "application/ld+json" {
					if n.FirstChild != nil {
						var jobLd map[string]interface{}
						err := json.Unmarshal([]byte(n.FirstChild.Data), &jobLd)
						if err == nil && jobLd["@type"] == "JobPosting" {
							job := parseJobPostingLD(jobLd)
							if job != nil && time.Since(job.PostedAt) < 7*24*time.Hour {
								jobs = append(jobs, *job)
							}
						}
					}
				}
			}
		}

		// Continue recursion
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			RecursiveFn(c)
		}
	}

	RecursiveFn(doc)
	return jobs
}

func extractTextByTag(n *html.Node, tag string) string {
	var result string
	var RecursiveFn func(*html.Node)
	RecursiveFn = func(n *html.Node) {
		if result != "" {
			return
		}
		if n.Type == html.ElementNode && n.Data == tag {
			if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
				result = n.FirstChild.Data
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			RecursiveFn(c)
		}
	}
	RecursiveFn(n)
	return result
}

func extractTextByClass(n *html.Node, tag string, class string) string {
	var result string
	var RecursiveFn func(*html.Node)
	RecursiveFn = func(n *html.Node) {
		if result != "" {
			return
		}
		if n.Type == html.ElementNode && n.Data == tag {
			for _, attr := range n.Attr {
				if attr.Key == "class" && strings.Contains(attr.Val, class) {
					if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
						result = n.FirstChild.Data
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			RecursiveFn(c)
		}
	}
	RecursiveFn(n)
	return result
}

func extractHrefByTag(n *html.Node, tag string) string {
	var result string
	var RecursiveFn func(*html.Node)
	RecursiveFn = func(n *html.Node) {
		if result != "" {
			return
		}
		if n.Type == html.ElementNode && n.Data == tag {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					result = "https://remoteok.com" + attr.Val
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			RecursiveFn(c)
		}
	}
	RecursiveFn(n)
	return result
}

func parseRelativeDate(text string) time.Time {
	text = strings.TrimSpace(text)
	now := time.Now()

	switch {
	case strings.HasSuffix(text, "d"):
		if d, err := strconv.Atoi(strings.TrimSuffix(text, "d")); err == nil {
			return now.AddDate(0, 0, -d)
		}
	case strings.HasSuffix(text, "h"):
		if h, err := strconv.Atoi(strings.TrimSuffix(text, "h")); err == nil {
			return now.Add(-time.Duration(h) * time.Hour)
		}
	case strings.HasSuffix(text, "m"):
		if m, err := strconv.Atoi(strings.TrimSuffix(text, "m")); err == nil {
			return now.Add(-time.Duration(m) * time.Minute)
		}
	}

	return now
}

func guessLocationFromTree(n *html.Node) string {
	var location string
	var search func(*html.Node)
	search = func(node *html.Node) {
		if location != "" {
			return
		}
		if node.Type == html.TextNode {
			if strings.Contains(strings.ToLower(node.Data), "remote") {
				location = strings.TrimSpace(node.Data)
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			search(c)
		}
	}
	search(n)
	return location
}

func generateUUID() string {
	return uuid.New().String()
}

func extractCity(loc string) string {
	parts := strings.Split(loc, ",")
	if len(parts) >= 1 {
		return strings.TrimSpace(parts[0])
	}
	return ""
}

func extractCountry(loc string) string {
	parts := strings.Split(loc, ",")
	if len(parts) >= 2 {
		return strings.TrimSpace(parts[1])
	}
	if strings.Contains(strings.ToLower(loc), "germany") {
		return "Germany"
	}
	if strings.Contains(strings.ToLower(loc), "usa") {
		return "USA"
	}
	return "Remote"
}

func parseJobPostingLD(data map[string]interface{}) *models.Job {
	title, _ := data["title"].(string)

	company := ""
	if org, ok := data["hiringOrganization"].(map[string]interface{}); ok {
		company, _ = org["name"].(string)
	}

	city := ""
	country := ""
	if locs, ok := data["jobLocation"].([]interface{}); ok && len(locs) > 0 {
		if loc, ok := locs[0].(map[string]interface{}); ok {
			if addr, ok := loc["address"].(map[string]interface{}); ok {
				city, _ = addr["addressLocality"].(string)
				country, _ = addr["addressCountry"].(string)
			}
		}
	}

	dateStr, _ := data["datePosted"].(string)
	postedAt := time.Now()
	if parsed, err := time.Parse(time.RFC3339, dateStr); err == nil {
		postedAt = parsed
	}

	logo := ""
	if org, ok := data["hiringOrganization"].(map[string]interface{}); ok {
		if logoObj, ok := org["logo"].(map[string]interface{}); ok {
			logo, _ = logoObj["url"].(string)
		}
	}

	link, _ := data["url"].(string)
	if link == "" {
		link = fmt.Sprintf("https://remoteok.com/remote-jobs/remote-%s", strings.ReplaceAll(strings.ToLower(title), " ", "-"))
	}

	return &models.Job{
		ID:        uuid.New().String(),
		Title:     strings.TrimSpace(title),
		Company:   strings.TrimSpace(company),
		City:      strings.TrimSpace(city),
		Country:   strings.TrimSpace(country),
		Link:      link,
		LogoURL:   logo,
		PostedAt:  postedAt,
		CreatedAt: time.Now(),
	}
}
