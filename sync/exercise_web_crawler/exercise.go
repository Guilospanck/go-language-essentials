package main

import (
	"fmt"
	"sync"
)

var cachedURLs sync.Map

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan Test, quit chan bool) {
	defer func() { quit <- true }()

	_, ok := cachedURLs.Load(url)
	if ok {
		return
	}
	cachedURLs.Store(url, url)

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	test := Test{
		body: body,
		url:  url,
	}
	ch <- test

	urlsToBeQuit := make([]chan bool, len(urls))
	for index, u := range urls {
		urlsToBeQuit[index] = make(chan bool)
		go Crawl(u, depth-1, fetcher, ch, urlsToBeQuit[index])
	}

	// wait for all child "quits" goroutines
	for _, childstatus := range urlsToBeQuit {
		<-childstatus
	}

}

type Test struct {
	url  string
	body string
}

func main() {
	ch := make(chan Test)
	quit := make(chan bool)
	defer close(ch)
	defer close(quit)
	go Crawl("https://golang.org/", 14, fetcher, ch, quit)

	for {
		select {
		case i := <-ch:
			fmt.Printf("found: %s %q\n", i.url, i.body)
		case <-quit:
			return
		}
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
