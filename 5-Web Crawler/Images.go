// find_in_page
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var document = MakeHttpRequest(1)
	// Get page count
	var allLIs = document.Find(".paginator").Find("ul li")
	var lastPage = document.Find(".paginator").Find("ul li").Eq(allLIs.Length() - 2)
	pageCount, _ := ToLatinDigits(lastPage.Find("a").Text())
	var nextPage int64 = 2
	for nextPage <= pageCount {
		nextPage++
		document = MakeHttpRequest(nextPage)

		// Find and print image URLs
		document.Find(".o-listView__itemInfo").Each(func(index int, element *goquery.Selection) {
			class, _ := element.Attr("class")

			fmt.Println(class, element.Text())
		})
	}

	// Find and print image URLs

	fmt.Println(pageCount)

}

func ToLatinDigits(persianNumber string) (i int64, err error) {
	var LatinDigits = strings.ReplaceAll(persianNumber, "۰", "0")
	LatinDigits = strings.ReplaceAll(LatinDigits, "۷", "7")
	return strconv.ParseInt(LatinDigits, 10, 64)

}

func MakeHttpRequest(pageNumber int64) *goquery.Document {
	// Make HTTP request
	req, _ := http.NewRequest("GET", "https://jobinja.ir/jobs?filters%5Bjob_categories%5D%5B%5D=%D9%88%D8%A8%D8%8C%E2%80%8C+%D8%A8%D8%B1%D9%86%D8%A7%D9%85%D9%87%E2%80%8C%D9%86%D9%88%DB%8C%D8%B3%DB%8C+%D9%88+%D9%86%D8%B1%D9%85%E2%80%8C%D8%A7%D9%81%D8%B2%D8%A7%D8%B1&filters%5Bkeywords%5D%5B0%5D=&filters%5Blocations%5D%5B%5D=%D8%AA%D9%87%D8%B1%D8%A7%D9%86&sort_by=relevance_desc&page="+strconv.FormatInt(pageNumber, 10), nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("cookie", "__cfduid=d3f05b23a1361bf6bca9e9609ed485b5d1572807113; _ga=GA1.2.925358084.1572807126; _gid=GA1.2.608524599.1572807126; logglytrackingsession=5f9314f7-4f02-4eab-83f1-6cd1c8901c9a; remember_82e5d2c56bdd0811318f0cf078b78bfc=eyJpdiI6Ik96NngyRTV1VlBPMmZjTTFnamZmU3c9PSIsInZhbHVlIjoiOTJMVGdaYTBMODkyUUowTnNWWkxxbks3eCtyQ0pqS0lJUVR5eWlIVit3ekNRSUVuNWNqYzRRaDY2MTNuOHl2KzhERTNrcmVtY2kxc1wvUUxvNERGb1ErZ2t6dVFcL0d2ZVwvRmNwS0o3NThEKzQ9IiwibWFjIjoiNWE0NTU1Nzk2ZDhlNjZhNWY0YzQ5ZDhlNGU4Y2IzYjU3NDVhNzdhYTc4OWMzYTgxMjExYmZmMDQzOWE3YTI2ZiJ9; device_id=eyJpdiI6IldYM253WmkybytwZTB1c0k5ZzYzUHc9PSIsInZhbHVlIjoid2ZwRG1vZFwvMnFmcUF5aTJIcHB1c3c9PSIsIm1hYyI6IjBmYzIzNTgzZjllYjhjOTEyODFmNjY4MjUwZDZjZWUzOTYwYzExMmU4ZGQzMzY4MDc1OTY0NzQ5ZmQ3ODY5YWQifQ%3D%3D; XSRF-TOKEN=eyJpdiI6IkU1dnV3TkJFaWlJeTA0UGtuZUw2UGc9PSIsInZhbHVlIjoid00rSDEyUVZwMG1MSVhOak43cE9jT3NzUXBQblJsdGRaQXFJd0wrb0dyNkNNNkVDU3ZcL1wvQmVvdkpVd256U0RJMzZEbzRwTFNia1NrcENXcEFJZkRxQT09IiwibWFjIjoiY2RmZmMyMTczMmY3ODNkZjIwYzc2ZDlmZTUyODgzZmU1OTAwODk4ZGVjYWViYThjNzMwZjQ4YzY3YjI0YjgwNCJ9; JSESSID=eyJpdiI6IjVtK3JRNjYweEREOEJTenhnc2pxaUE9PSIsInZhbHVlIjoiVU5VUzNTV3V3bmtaT3h6V0dWRE95ck1adFl6T25TdEZBNWZBbkVrSXFpSURRelBORXkyWk9nc0huXC9veE00THUwWCtIeGVPSm1FeDNDc0k4WWhPdXNRPT0iLCJtYWMiOiJjMzQ0ZjI2MjNjZTFhZDI4MDk0Y2I5ZDRhODJkMzk0YTg3YWI4MTg4Y2NmNjhkZjEzYzBkZGUyM2RhZmVjNWI5In0%3D; user_mode=eyJpdiI6IjFGNUtOOTkzeFp2UlJJRnFpV1c4ZEE9PSIsInZhbHVlIjoiZnJxSm14RW4zYkxweTlsRUx2aGFIQT09IiwibWFjIjoiMGExZTc0Nzc5M2JhN2M4MzM4MGU1NDRmY2E0NWRiODczNTgzM2JhNjQ4YTA5MzY2YzA5NTk3NTAwMTg5NmEwZCJ9; _gat_gtag_UA_129125700_1=1")
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, _ := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	return document
}
