package pkg

import (
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

func GetPinterestUrls(searchTerm string) []string {
	c := colly.NewCollector()
	var pinterestUrls []string

	c.OnHTML("#main > div > div > div > a", func(e *colly.HTMLElement) {
		href := strings.Replace(e.Attr("href"), "/url?q=", "", 1)
		if strings.Contains(href, "pinterest.com") {
			pinterestUrls = append(pinterestUrls, href)
		}
	})
	c.Visit("http://www.google.co.in/search?hl=en&num=100&q=" + url.QueryEscape(searchTerm))

	return pinterestUrls
}
