package scrap

import (
	"fmt"
	"time"
	m "github.com/ahnsv/blog-scrapper/models"
	"github.com/gocolly/colly"
)

var websites = make([]string, 0)
type Post = *m.Post
func AddWebsite(s string) {
	websites = append(websites, s)
}

func Init() {
	posts := []Post
	c := colly.NewCollector(
		colly.AllowedDomains("https://taegon.kim/archives/category/tiptech"),
		colly.Async(true),
	)	
	c.OnHTML("#main article", func(e *colly.HTMLElement) {
		temp := Post{}
		temp.Title := e.ChildText(".entry-title a")
		temp.Content := e.ChildAttr(".entry-title a", "href")
		temp.Date := time.Now()
		temp.Tags := e.ChildText(".entry-category span a")
		posts = append(posts, temp)
	} )
}