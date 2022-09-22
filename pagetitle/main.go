package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func fetchHTML(URL string) (io.ReadCloser, error) {

	//GET the url
	resp, err := http.Get(URL)

	var errOutput error

	if err != nil {
		//.Fatalf() prits the error nd exits the process
		//log.Fatalf("error fetching URL: %v\n", err)
		errOutput = fmt.Errorf("error fetching URL: %v\n", err)

	}
	//close the response body regardless of the outcome
	//readerCloser := resp.Body.Close()

	//check if response status code is ok, if not report error
	if resp.StatusCode != http.StatusOK {
		errOutput = fmt.Errorf("response status code was %d\n", resp.StatusCode)
	}

	//check reponse content type
	ctype := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(ctype, "text/html") {
		errOutput = fmt.Errorf("response content type was %s not text/html\n", ctype)
	}

	return resp.Body, errOutput
}

func extractTitle(body io.ReadCloser) (string, error) {
	tokenizer := html.NewTokenizer(body)

	var title string
	var errorOutput error

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			errorOutput = fmt.Errorf("error tokenizing HTML: %v", tokenizer.Err())
		}

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if "title" == token.Data {
				tokenType = tokenizer.Next()

				if tokenType == html.TextToken {
					title = tokenizer.Token().Data
					break
				}

			}
		}
	}
	return title, errorOutput
}

func fetchTitle(URL string) (string, error) {
	body, err := fetchHTML(URL)
	//if err != nil {
	//	return "theres error", err
	//}
	title, err := extractTitle(body)
	//if err != nil {
	//	return "theres error", err
	//}
	return title, err
}

func main() {
	//if the caller didn't provide a URL to fetch, then:
	if len(os.Args) < 2 {
		//print the usage and exist with eror
		fmt.Printf("usage:\n pagetitle <url>\n")
		os.Exit(1)
	}

	title, err := fetchTitle(os.Args[1])
	if err != nil {
		log.Fatalf("error fetching page title: %v\n", err)
	}
	fmt.Println(title)
}
