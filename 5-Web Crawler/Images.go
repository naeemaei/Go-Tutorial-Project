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

var (
	baseURL string    = "https://jobinja.ir/"
	jobPage string    = "jobs?"
	params  [4]string = [4]string{
		"filters[job_categories][0]=وب،‌ برنامه‌نویسی و نرم‌افزار",
		"filters[locations][0]=تهران",
		"sort_by=relevance_desc",
		"page=",
	}
)

func main() {
	TestConnection()

	var document = MakeHttpRequest(baseURL+jobPage+strings.Join(params[:], "&"), 1)
	// Get page count
	var allLIs = document.Find(".paginator").Find("ul li")
	var lastPage = document.Find(".paginator").Find("ul li").Eq(allLIs.Length() - 2)
	pageCount, _ := ToLatinDigits(lastPage.Find("a").Text())
	var nextPage int64 = 2
	//pageCount = 2
	for nextPage <= pageCount {

		// Find and print image URLs
		document.Find(".o-listView__itemInfo").Each(func(index int, element *goquery.Selection) {

			// Get company name
			var company = element.Find(".c-icon--construction").Parent().Find("span").Text()
			fmt.Println(company)

			// Get city name
			var place = element.Find(".c-icon--place").Parent().Find("span").Text()
			fmt.Println(place)

			// Get job title name
			var jobTitle = element.Find(".c-jobListView__titleLink").Text()
			var jobLink, _ = element.Find(".c-jobListView__titleLink").Attr("href")
			// fmt.Println(jobLink)

			// Go to detail page
			// Save in db
			condb := GetConnection()
			var jobId, _ = CreateMasterRecord(condb, jobTitle, company, place)

			var newDocument = MakeHttpRequest(jobLink, 0)
			newDocument.Find(".c-infoBox__item").Each(func(index int, element *goquery.Selection) {
				fmt.Println(element.Find("h4").Text())
				fmt.Println(element.Find("span").Text())
				fmt.Println(newDocument.Find(".s-jobDesc ").Text())
				CreateDetailRecord(condb, int(jobId), element.Find("h4").Text(), element.Find("span").Text())
			})

			CreateDetailRecord(condb, int(jobId), "Description", newDocument.Find(".s-jobDesc ").Text())

		})
		nextPage++
		document = MakeHttpRequest(baseURL+jobPage+strings.Join(params[:], "&"), nextPage)
	}

	// Find and print image URLs

	fmt.Println(pageCount)

}

func MakeHttpRequest(pageAddress string, pageNumber int64) *goquery.Document {
	// Make HTTP request

	if pageNumber > 0 {
		pageAddress += strconv.FormatInt(pageNumber, 10)
	}
	req, _ := http.NewRequest("GET", pageAddress, nil)
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

func ToLatinDigits(persianNumber string) (i int64, err error) {
	var LatinDigits = strings.ReplaceAll(persianNumber, "۰", "0")
	LatinDigits = strings.ReplaceAll(LatinDigits, "۷", "7")
	LatinDigits = strings.ReplaceAll(LatinDigits, "۱", "1")
	LatinDigits = strings.ReplaceAll(LatinDigits, "۲", "2")
	LatinDigits = strings.ReplaceAll(LatinDigits, "۳", "3")
	LatinDigits = strings.ReplaceAll(LatinDigits, "۴", "4")
	LatinDigits = strings.ReplaceAll(LatinDigits, "۵", "5")
	LatinDigits = strings.ReplaceAll(LatinDigits, "۶", "6")
	LatinDigits = strings.ReplaceAll(LatinDigits, "۸", "8")
	LatinDigits = strings.ReplaceAll(LatinDigits, "۹", "9")

	return strconv.ParseInt(LatinDigits, 10, 64)

}
