package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// fetchDataFromBlic scrapes news items from the blic.rs.
func fetchDataFromBlic(url string) []map[string]string {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Create a GoQuery document from the HTTP response
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Create a slice of maps. Each map will hold the data for one news item.
	var newsItems []map[string]string

	// Find elements with the class 'news__content' and print their contents
	doc.Find(".news__content").Each(func(i int, s *goquery.Selection) {

		newsItem := map[string]string{}

		// Find the <h2> element, then find the <a> element within it
		a := s.Find("h2 a")
		href, exists := a.Attr("href")
		if !exists {
			return
		}

		href = strings.TrimSpace(href)
		newsItem["link"] = href

		// Get the text content of the <a> element
		text := a.Text()
		text = strings.TrimSpace(text)

		newsItem["headline"] = text

		// Find and extract the date from the <time> element
		date := s.Find("time").Text()
		date = strings.TrimSpace(date)
		newsItem["date"] = date

		// Append the newly created map to the slice
		newsItems = append(newsItems, newsItem)

	})
	return newsItems
}

// fetchDataFromMondo scrapes news items from the mondo.rs.
func fetchDataFromMondo(url string) []map[string]string {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Create a GoQuery document from the HTTP response
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Create a slice of maps. Each map will hold the data for one news item.
	var newsItems []map[string]string

	// Find elements with the class 'news-wrapper' and print their contents
	doc.Find("article.news-wrapper").Each(func(i int, s *goquery.Selection) {

		newsItem := map[string]string{}

		// Extract the article link and title
		a := s.Find("a.title-wrapper")
		href, exists := a.Attr("href")
		if !exists {
			return
		}
		newsItem["link"] = strings.TrimSpace(href)

		title := s.Find("h2.title").Text()
		newsItem["headline"] = strings.TrimSpace(title)

		date := s.Find("p.time").Text()
		date = strings.TrimSpace(date)              // Remove leading and trailing whitespace
		date = strings.Replace(date, "|\n", "", -1) // Remove the specific pattern "|\n"
		newsItem["date"] = date

		// Append the newly created map to the slice
		newsItems = append(newsItems, newsItem)

	})
	return newsItems
}

// saveToJson serializes the given slice to JSON and saves it to a file.
func saveToJson(newsItems []map[string]string, filePath string) {
	jsonData, err := json.MarshalIndent(newsItems, "", "    ")
	if err != nil {
		log.Fatal("Error marshalling JSON: ", err)
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		log.Fatal("Error writing JSON to file: ", err)
	}

	fmt.Printf("News items have been saved to %s\n", filePath)
}

// combineSlices combines two slices of map[string]string into one.
func combineSlices(slice1, slice2 []map[string]string) []map[string]string {
	return append(slice1, slice2...)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <search term>")
		return
	}

	// Combine the command-line arguments into a search term
	searchTerm := strings.Join(os.Args[1:], " ")

	// URL encode the search term to ensure it's safe to use in a URL
	encodedSearchTerm := url.QueryEscape(searchTerm)

	blicUrl := "https://www.blic.rs/search?q=" + encodedSearchTerm
	mondoUrl := "https://mondo.rs/search/1/1?q=" + encodedSearchTerm

	newsArticlesBlic := fetchDataFromBlic(blicUrl)
	newsArticlesMondo := fetchDataFromMondo(mondoUrl)

	combinedNewsItems := combineSlices(newsArticlesBlic, newsArticlesMondo)

	saveToJson(combinedNewsItems, "./combined_news_items.json")
}
