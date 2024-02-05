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

type MenuItem struct {
	Date        string
	Foods       []string
	CalorieText string
}

func scrape(menu *[]MenuItem) {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://sks.iuc.edu.tr/tr/yemeklistesi`),
		chromedp.WaitVisible(`#tab-ogle`, chromedp.ByQuery),
		chromedp.OuterHTML(`html`, &htmlContent, chromedp.ByQuery),
	)

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#tab-ogle section.monu-container > table > tbody > tr > td > b").Each(func(i int, s *goquery.Selection) {
		date := s.Text()
		*menu = append(*menu, MenuItem{Date: date})
	})

	doc.Find("#tab-ogle section.monu-container > table > tbody > tr > td.monu").Each(func(i int, s *goquery.Selection) {
		food := s.Text()
		foodArray := strings.Split(food, "\n")
		var tempArray []string

		for _, v := range foodArray {
			if v != "" {
				tempArray = append(tempArray, v)
			}
		}

		(*menu)[i].Foods = tempArray
	})

	doc.Find("#tab-ogle section.monu-container > table > tbody > tr:nth-child(3) > td").Each(func(i int, s *goquery.Selection) {
		calorieText := s.Text()

		if calorieText != "" {
			(*menu)[i].CalorieText = calorieText
		}
	})
}

func printMenu(menu []MenuItem) {
	for _, menuItem := range menu {
		fmt.Printf("%s\n", menuItem.Date)
		fmt.Println("--------------------")
		for _, food := range menuItem.Foods {
			fmt.Printf("-- %s\n", food)
		}
		fmt.Println("--------------------")
		fmt.Printf("%s\n", menuItem.CalorieText)
		fmt.Println()
		fmt.Println("==========================")
		fmt.Println()
	}
}

func main() {
	var menu []MenuItem

	scrape(&menu)

	printMenu(menu)
}
