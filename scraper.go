package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly"
)

// basic structure of info I want of the macbooks
type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {
	c := colly.NewCollector()
	c.AllowedDomains = []string{"www.bestbuy.com"}

	// mimics web browser to avoid bot detection

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	})

	var items []item

	// finds attributes/text listed on site

	c.OnHTML("li.sku-item", func(h *colly.HTMLElement) {
		price := h.ChildText("div.priceView-hero-price span[aria-hidden=true]")
		item := item{
			Name:   h.ChildText("h4.sku-title"),
			Price:  price,
			ImgUrl: h.ChildAttr("img", "src"),
		}

		items = append(items, item)

		fmt.Println(items)
	})

	// looks for next page button on page

	c.OnHTML("[class=sku-list-page-next] a", func(h *colly.HTMLElement) {
		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(next_page)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	startURL := "https://www.bestbuy.com/site/searchpage.jsp?st=macbook&_dyncharset=UTF-8&_dynSessConf=&id=pcat17071&type=page&sc=Global&cp=1&nrp=&sp=&qp=&list=n&af=true&iht=y&usc=All+Categories&ks=960&keys=keys"
	err := c.Visit(startURL)
	if err != nil {
		fmt.Println("Error:", err)
	}

	//allows page to load to avoid scroll loading

	time.Sleep(5 * time.Second)

	//encodes data for JSON and inputs indents and spacing for clearer reading

	exportedData, err := json.MarshalIndent(items, "", "    ")
	if err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}

	//writes the json file through writeJSON func and names file

	filename := "bestbuy_macbook_data.json"
	err = writeJSON(filename, exportedData)
	if err != nil {
		fmt.Println("Error writing JSON:", err)
	} else {
		fmt.Println("Data exported to", filename)
	}
}

func writeJSON(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	return err
}
