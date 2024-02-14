package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/common-nighthawk/go-figure"
	"github.com/gocolly/colly/v2"
)

func main() {

	asciiArt :=
		`                                                                                                                                        
                  
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠿⣿⣿⠛⠛⣛⣿⣿⣛⠛⢻⡿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠿⢛⣫⡅⢠⣾⣿⣿⣿⠀⣿⣿⣿⡏⣠⠸⣿⠆⣶⣾⣿⣿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣶⡀⠟⡀⠻⢿⣿⣿⣿⠀⣿⣿⣿⠁⣿⡆⢿⢠⣿⡿⢋⣴⣶⣦⠙⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⡿⠿⣿⣿⣿⣿⣿⣷⠀⣿⣷⠆⣈⣭⣭⣤⣬⣽⣤⣤⣭⣿⡀⣸⣿⠃⣾⣿⣿⣿⣰⣿⣿⡛⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣷⡐⣦⡍⠛⠿⠿⢛⣡⣴⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⡙⠻⠛⣩⣿⣿⠟⣡⣬⣽⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣷⠈⣴⣿⣂⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠋⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡛⠛⣡⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⢨⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⣱⣷⠀⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⣽⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠋⣭⣭⣴⣿⣿⠈⣦⣙⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠛⠛⣸⣿⠟⠙⣿⣿⡄⢿⣿⣷⣦⠹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⢠⣿⣿⣿⠟⣰⡄⣿⣿⣿⣦⠻⣿⣿⡈⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⣼⣿⠇⣴⣶⣿⣿⡏⢹⣿⣿⣧⣿⣿⣷⣬⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣯⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣿⣿⣧⣤⣤⣤⣤⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⠙⣿⣿⡏⢸⣿⣿⣇⢸⣿⣿⡿⢸⣿⣿⠏⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣏⣭⣽⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡌⢻⣧⠘⢿⣿⣿⡘⣿⡏⣴⡿⠛⢋⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣟⠟⣻⣿⣿⣿
⣿⣿⣿⣷⠿⡧⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⢸⣿⣿⡆⣿⣿⣷⡘⢀⣿⡇⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣴⡖⣾⣿⣿⣿
⣿⣿⣿⣿⡟⣛⠝⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣶⣦⢸⡇⢸⡟⢻⣷⣾⣿⢁⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⢧⣄⣲⣿⣿⣿⣿
⣿⣿⣿⣿⣿⡖⡛⡘⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡀⢿⠈⠃⡌⠿⣿⠃⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣡⣔⢣⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣞⣏⢮⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣄⣰⣿⣶⣦⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⣯⣾⣹⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣧⡿⠕⡹⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⣅⣄⢕⣽⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⣫⡔⣹⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⢟⢓⠔⣽⣽⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣮⣵⠊⡞⢍⣻⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⢻⡝⣎⢼⣣⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⣽⣡⠽⡷⣼⣿⢭⣟⡟⣻⢻⠛⢟⣏⣿⣇⡿⢣⣪⣯⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣿⣯⣼⣼⣦⣿⣼⣿⣿⣬⣽⣾⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
                                                         
                                                                                                           
	`
	fmt.Println(asciiArt)

	myFigure := figure.NewFigure("AKINCI - WEB - SCRAPPER V3", "", true)
	myFigure.Print()

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
			htmlContent := e.Text
			fmt.Println(htmlContent)
			writeToFile("websitesi.html", htmlContent)
			log.Printf("HTML içeriği çekildi ve 'html_output.html' dosyasına kaydedildi.\n")
		})
	}

	if scrapeLinks {
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			fmt.Println(link)
			writeToFile("links.txt", link+"\n")
			log.Printf("Linkler çekildi ve 'links_output.txt' dosyasına kaydedildi.\n")
		})
	}

	if captureScreenshot {
		if err := captureScreen(targetURL); err != nil {
			log.Fatalf("Ekran görüntüsü alınırken hata oluştu: %v", err)
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

func captureScreen(url string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, chromedp.Navigate(url), chromedp.CaptureScreenshot(&buf)); err != nil {
		return err
	}

	fileName := "screenshot.png"
	if err := ioutil.WriteFile(fileName, buf, 0644); err != nil {
		return err
	}
	fmt.Println("Ekran görüntüsü başarıyla kaydedildi:", fileName)
	return nil
}

func writeToFile(filename string, data string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Dosya açılırken hata oluştu: %v", err)
	}
	defer f.Close()

	if _, err := f.WriteString(data); err != nil {
		log.Fatalf("Dosyaya yazılırken hata oluştu: %v", err)
	}
}
