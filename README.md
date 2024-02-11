# Go News Scraper

The Go News Scraper is a dynamic command-line application that aggregates news articles based on user-provided search terms. Built with Go and leveraging the GoQuery package, it simplifies the process of fetching and parsing web content, demonstrating an efficient approach to web scraping and data aggregation.

## Features

- **Dynamic Search**: Supports user-defined search terms to dynamically scrape related news articles.
- **Multiple Sources**: Capable of scraping from various predefined news websites, with easy extensibility for more sources.
- **JSON Output**: Compiles and saves scraped article data into a structured JSON file, facilitating further data analysis or usage.

## Getting Started

### Prerequisites

- Go (1.16 or newer): Ensure Go is installed on your system by running `go version` in your terminal. If you need to install Go, follow the instructions on the [official Go website](https://golang.org/dl/).

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/acast83/go_web_scraper.git
cd go_web_scraper
```
### Usage

Run the scraper using `go run`, followed by the search term(s) you wish to scrape articles for:

```bash
go run scraper.go "your search term here"
```

This will execute the scraper and save the aggregated news items into a file named combined_news_items.json in your current directory.

### Building the Application (Optional)
To compile the application into an executable binary, use:

```bash
go build -o news_scraper
```
This creates an executable named news_scraper that you can run directly:

```bash
./news_scraper "your search term here"
```



