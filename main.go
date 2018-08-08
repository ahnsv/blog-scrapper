package main

import (
	. "fmt"
	"github.com/ahnsv/blog-scrapper/cache"
	"github.com/ahnsv/blog-scrapper/scrap"
)

func main() {
	Println("Scrap Initiating...")
	posts := scrap.Init()
	Println("Posts are ", posts)
	cache.ClientInit()
}
