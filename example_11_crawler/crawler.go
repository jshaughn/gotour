package main

import (
	"fmt"
	"sync"
	"time"
)

type UrlCache struct {
	v   map[string][]string
	mux sync.Mutex
}

func (c *UrlCache) Put(key string, value []string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key] = value
	c.mux.Unlock()
}

func (c *UrlCache) PutIfNotExists(key string, value []string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	_, ok := c.v[key]
	if !ok {
		c.v[key] = value
	}
	return !ok
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c UrlCache) {
	if depth <= 0 {
		return
	}
	if c.PutIfNotExists(url, nil) {
		fmt.Printf("Crawl!: %s\n", url)
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Found: %s %q\n", urls, body)
		c.Put(url, urls)
		for _, u := range urls {
			go Crawl(u, depth-1, fetcher, c)
		}
		return
	}
}

func main() {
	c := UrlCache{v: make(map[string][]string)}
	go Crawl("http://golang.org/", 4, fetcher, c)
	time.Sleep(20 * time.Second)
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

