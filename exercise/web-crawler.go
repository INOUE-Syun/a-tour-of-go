package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type SafeCache struct {
	fetched map[string]struct{} // Url is cachekey
	mux     sync.Mutex
}

func (sc *SafeCache) AddCache(k string) {
	sc.mux.Lock()
	sc.fetched[k] = struct{}{}
	sc.mux.Unlock()
}

func (sc *SafeCache) HasCache(k string) bool {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	_, has := sc.fetched[k]
	return has
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	cache := &SafeCache{fetched: make(map[string]struct{})}

	var wg sync.WaitGroup
	var worker func(string, int)

	worker = func(url string, depth int) {
		if cache.HasCache(url) || depth <= 0 {
			return
		}
		cache.AddCache(url)

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)

		wg.Add(len(urls))
		for _, url := range urls {
			go func(u string) {
				worker(u, depth-1)
				wg.Done()
			}(url)
		}
	}

	worker(url, depth)
	wg.Wait()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
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
