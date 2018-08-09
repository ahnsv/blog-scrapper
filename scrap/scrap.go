package scrap

import (
	// "flag"
	"fmt"
	m "github.com/ahnsv/blog-scrapper/models"
	"github.com/gocolly/colly"
	"log"
	// "os"
	"time"
)

var blogs = make([]string, 0)
var links = make([]string, 0)

type Post m.Post
type Content m.Content

func AddWebsite(s string) {
	blogs = append(blogs, s)
}

func Init() (posts []Post, links []string) {
	// var url string
	// flag.StringVar(&url, "URL", "", "URL to crawl")
	// flag.Parse()
	// if url == "" {
	// 	log.Println("url is required")
	// 	os.Exit(1)
	// }

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
		links = append(links, temp.Content)
		// fmt.Printf("here are contents, just for checking : %v", Contents(links, e))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://taegon.kim/archives/category/tiptech")
	for i := 0; i < 100; i++ {
		err := c.Visit("https://taegon.kim/archives/category/tiptech/page/" + i)
		if err != nil {
			break
		}
	}

	c.Wait()
	return posts, links
}
