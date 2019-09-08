 // Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 125.

// Findlinks2 does an HTTP GET on each URL, parses the
// result as HTML, and prints the links within it.
//
// Usage:
//	findlinks url ...
package main

 import (
	 "fmt"
	 "golang.org/x/net/html"
	 "net/http"
	 "os"
 )

// visit appends to links each link found in n, and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

//!+
func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
/********************************************************
	func findLinksLog(url string) ([]string, error) {
		log.Printf("findLinks %s", url)
		return findLinks(url)
	}

	log.PrintIn(findLinks(url))
	links, err := findLinks(url)
	log.PrintIn(links, err)

	func Size(rect image.Rectangle) (width, height int)
	func Split(path string) (dir, file string)
	func HourMinSec(t time.Time) (hour, minute, second int)

	// CountWordsAndlmages выполняет HTTP-запрос GET HTML-документа
	// url и возвращает количество слов и изображений в нем.
	func CountWordsAndImages(url string) (words, images int, err error) {
		resp, err := http.Get(url)
		if err != nil {
			return
		}
		doc, err := html.Parse(resp.Body)
		resp.Body .CloseQ
		if err != nil {
			err = fmt.Errorf("parsing HTML: %s", err)
			return
		}
		words, images = countWordsAndlmages(doc)
		return
	}
	func countWordsAndImages(n *html.Node)(words,images int){ / * . . . * / }

	value, ok := cache.Lookup(key)
	if !ok {

				// ...cache[key] не существует...
	}

	********************************************************/
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

//!-