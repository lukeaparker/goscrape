package main

import (
	"fmt"

	"log"

	"github.com/gocolly/colly"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Product gbygyg
type Product struct {
	gorm.Model
	Price uint
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Instantia
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("#main > div > div.col.large-9 > div > div.products.row.row-small.large-columns-3.medium-columns-3.small-columns-1.has-shadow.row-box-shadow-1.row-box-shadow-3-hover.has-equal-box-heights > div.product-small.col.has-hover.out-of-stock.product.type-product.post-287.status-publish.first.outofstock.product_cat-boxzy.product_tag-cnc-mill.has-post-thumbnail.taxable.shipping-taxable.purchasable.product-type-variable.has-default-attributes > div > div.product-small.box > div.box-text.box-text-products.flex-row.align-top.grid-style-3.flex-wrap > div.price-wrapper > span > span", func(e *colly.HTMLElement) {
		link := e.Attr("p")

		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Create
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://boxzy.com/product-category/boxzy/")
}
