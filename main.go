package main

import (
	. "fmt"
	"github.com/ahnsv/blog-scrapper/cache"
	"github.com/ahnsv/blog-scrapper/router"
	"github.com/ahnsv/blog-scrapper/scrap"
)

func main() {
	Println("Scrap Initiating...")
	posts, links := scrap.Init()
	Println("Posts are ", posts)
	cache.ClientInit()
	router.Init()
}
