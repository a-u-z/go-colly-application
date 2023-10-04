package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// 規劃

// 一定會有的要素： 1. 造訪網站 2. 想要搜尋的字串（要可以是多個） 3.

// 需要記憶功能，傳過的資訊就不用再傳送依次
func Co(url string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.ptt.cc", "www.ptt.cc"),
	)
	extensions.RandomUserAgent(c)

	detailCollector := c.Clone()
	goQuery := WriteStringByBuilder(500, "div[class=r-ent] div[class=title] a:contains('[販售]')") // 販售 徵求

	c.OnHTML(goQuery, func(e *colly.HTMLElement) {
		lowerCaseE := strings.ToLower(e.Text)
		log.Printf("here is lowerCaseE:%+v", lowerCaseE)
		if (strings.Contains(lowerCaseE, "m1")) && strings.Contains(e.Text, "book") || strings.Contains(e.Text, "iphone") {
			baseUrl := "https://www.ptt.cc"
			subLink := e.Attr("href")

			detailCollector.Visit(WriteStringByBuilder(200, baseUrl, subLink))
		}
	})
	detailCollector.OnHTML("div[id=main-content]", func(e *colly.HTMLElement) {
		// log.Printf("here is e.Text:%+v", e.Text)
		price := refinePrice(e.Text[strings.Index(e.Text, "[售價]"):strings.Index(e.Text, "[交易方式/地點]")])
		priceInt, err := strconv.Atoi(price)
		if err != nil {
			log.Printf("here is err:%+v", err)
			return
		}
		if priceInt < 200000 {
			// 發送訊息給 telegram
			log.Printf("here is :%+v", refinePrice(e.Text[strings.Index(e.Text, "[售價]"):strings.Index(e.Text, "[交易方式/地點]")]))
		}
		// log.Printf("here is e.Text:%+v", len(e.Text))

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
		log.Println("visiting", r.URL.String())
	})

	c.Visit(url)
}
