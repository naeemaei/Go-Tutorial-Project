// find_images_in_page.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	// Make HTTP request
	response, err := http.Get("https://jobinja.ir/jobs?filters%5Bkeywords%5D%5B0%5D=&filters%5Bjob_categories%5D%5B%5D=%D9%88%D8%A8%D8%8C%E2%80%8C+%D8%A8%D8%B1%D9%86%D8%A7%D9%85%D9%87%E2%80%8C%D9%86%D9%88%DB%8C%D8%B3%DB%8C+%D9%88+%D9%86%D8%B1%D9%85%E2%80%8C%D8%A7%D9%81%D8%B2%D8%A7%D8%B1&filters%5Blocations%5D%5B%5D=%D8%AA%D9%87%D8%B1%D8%A7%D9%86&sort_by=published_at_desc&page=1")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find and print image URLs
	document.Find(".o-listView__itemTitle").Each(func(index int, element *goquery.Selection) {
		class, _ := element.Attr("class")

		fmt.Println(class, element.Text())
	})

}
