package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	var targetURL string
	var scrapeHTML, scrapeLinks bool

	flag.StringVar(&targetURL, "url", "", "Web sitesinin URL'si")
	flag.BoolVar(&scrapeHTML, "html", false, "HTML içeriğini çek")
	flag.BoolVar(&scrapeLinks, "links", false, "Linkleri çek")
	flag.Parse()

	if targetURL == "" {
		fmt.Println("Hata: URL belirtilmedi.")
		os.Exit(1)
	}

	c := colly.NewCollector()

	if scrapeHTML {
		c.OnHTML("html", func(e *colly.HTMLElement) {
			fmt.Println(e.Text)
		})
	}

	if scrapeLinks {
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			fmt.Println(link)
		})
	}

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Hata:", err)
	})

	err := c.Visit(targetURL)
	if err != nil {
		fmt.Println("URL'yi ziyaret ederken bir hata oluştu:", err)
		os.Exit(1)
	}
}
