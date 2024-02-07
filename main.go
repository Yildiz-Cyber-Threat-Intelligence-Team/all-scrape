package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
)

func main() {
	var targetURL string
	var scrapeHTML, scrapeLinks, captureScreenshot bool

	flag.StringVar(&targetURL, "url", "", "Web sitesinin URL'si")
	flag.BoolVar(&scrapeHTML, "html", false, "HTML içeriğini çek")
	flag.BoolVar(&scrapeLinks, "links", false, "Linkleri çek")
	flag.BoolVar(&captureScreenshot, "screenshot", false, "Sayfa görüntüsü al")
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

	if captureScreenshot {
		msl, fln := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
		defer fln()

		var screenshot []byte
		if err := chromedp.Run(msl, captureScreenshotTask(targetURL, &screenshot)); err != nil {
			log.Fatal(err)
		}

		screenshotFilename := "screenshot.png"
		if err := ioutil.WriteFile(screenshotFilename, screenshot, 0644); err != nil {
			log.Fatal(err)
		}
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

func captureScreenshotTask(url string, screenshot *[]byte) chromedp.Tasks {
	var err error
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			headers := map[string]interface{}{
				"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.9999.999 Safari/537.36",
			}
			err = chromedp.ActionFunc(func(ctx context.Context) error {
				return network.Enable().Do(ctx)
			}).Do(ctx)
			if err != nil {
				return err
			}
			err = chromedp.ActionFunc(func(ctx context.Context) error {
				return network.SetExtraHTTPHeaders(headers).Do(ctx)
			}).Do(ctx)
			if err != nil {
				return err
			}

			err = chromedp.Sleep(90 * time.Second).Do(ctx)
			if err != nil {
				return err
			}

			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			*screenshot, err = page.CaptureScreenshot().WithQuality(90).Do(ctx)
			return err
		}),
	}
}
