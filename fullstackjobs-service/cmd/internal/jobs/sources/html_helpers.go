package sources

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

func parseJobsFromHTML(htmlStr string) {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	var RecursiveFn func(*html.Node)
	// walks the DOM tree, starting from a node n
	RecursiveFn = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "tr" {
			// Check if class="job"
			for _, attr := range n.Attr {
				if attr.Key == "class" && strings.Contains(attr.Val, "job") {
					title := extractTextByTag(n, "h2")
					company := extractTextByTag(n, "h3")
					location := extractTextByClass(n, "span", "location")
					date := extractTextByTag(n, "time")

					if title != "" && company != "" {
						log.Printf("Job Title: %s, Company: %s, Location: %s, Date Posted: %s\n",
							strings.TrimSpace(title), strings.TrimSpace(company), strings.TrimSpace(location), strings.TrimSpace(date))
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			RecursiveFn(c)
		}
	}
	RecursiveFn(doc)
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
