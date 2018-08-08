package scrap

import (
	"fmt"
	m "github.com/ahnsv/blog-scrapper/models"
	"github.com/gocolly/colly"
	"time"
)

var blogs = make([]string, 0)

type Post m.Post

func AddWebsite(s string) {
	blogs = append(blogs, s)
}

func Init() []Post {
	posts := make([]Post, 0)
	c := colly.NewCollector(
		colly.AllowedDomains("taegon.kim"),
		colly.Async(true),
	)
	c.OnHTML("#main article", func(e *colly.HTMLElement) {
		temp := &Post{}
		temp.Title = e.ChildText(".entry-title a")
		temp.Content = e.ChildAttr(".entry-title a", "href")
		temp.Date = time.Now()
		temp.Tags = append(temp.Tags, e.ChildText(".entry-category span a"))
		posts = append(posts, *temp)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://taegon.kim/archives/category/tiptech")

	c.Wait()
	return posts
}
