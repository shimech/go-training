package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

func main() {
	baseUrl := "https://www.alexa.com"
	url := fmt.Sprintf("%s/topsites", baseUrl)
	response, _ := http.Get(url)

	buffer, _ := io.ReadAll(response.Body)

	detector := chardet.NewTextDetector()
	detectResult, _ := detector.DetectBest(buffer)
	fmt.Printf("charset: %s\n", detectResult.Charset)

	bufferReader := bytes.NewReader(buffer)
	reader, _ := charset.NewReaderLabel(detectResult.Charset, bufferReader)

	document, _ := goquery.NewDocumentFromReader(reader)

	descriptionCells := document.Find("div.DescriptionCell")
	urls := []string{}
	descriptionCells.Each(func(index int, descriptionCell *goquery.Selection) {
		href := descriptionCell.Find("a").AttrOr("href", "")
		url := fmt.Sprintf("%s%s\n", baseUrl, href)
		urls = append(urls, url)
	})

	filepath := "../txt/urls.txt"
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		file.Close()
		os.Exit(1)
	}
	for _, url := range urls {
		b := []byte(url)
		_, err := file.Write(b)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			file.Close()
			os.Exit(1)
		}
	}
	return
}
