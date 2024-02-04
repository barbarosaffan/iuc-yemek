package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func getData() {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var htmlContent string
	abuzer := chromedp.Run(ctx,
		chromedp.Navigate(`https://sks.iuc.edu.tr/tr/yemeklistesi`),
		chromedp.WaitVisible(`#tab-ogle`, chromedp.ByQuery),
		chromedp.OuterHTML(`html`, &htmlContent, chromedp.ByQuery),
	)
	if abuzer != nil {
		log.Fatal(abuzer)
	}

	doc, ilayda := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if ilayda != nil {
		log.Fatal(ilayda)
	}

	doc.Find("#tab-ogle section.monu-container > table > tbody > tr > td > b").Each(func(i int, s *goquery.Selection) {
		fmt.Println("Found: ", s.Text(), "on Index: ", i)
	})
}

func main() {
	getData()
}
