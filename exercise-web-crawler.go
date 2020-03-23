// reference to solution from https://github.com/golang/tour/blob/master/solutions/webcrawler.go

package main

import (
	"errors"
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

var fetched = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

var loading = errors.New("url load in progress") // sentinel value

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。 Done
	// TODO: 不重复抓取页面。 Done
	// 下面并没有实现上面两种情况：
	if depth <= 0 {
		fmt.Printf("<- Done with %v, depth 0.\n", url)
		return
	}
	fetched.Lock()
	if _, ok := fetched.m[url]; ok {
		fetched.Unlock()
		fmt.Printf("Already fetched: %v\n", url)
		return
	}
	//We mark the url to be loading to avoid others reloading it at the same time
	fetched.m[url] = loading
	fetched.Unlock()

	//We load it concurrently
	fmt.Printf("Fetch: %v\n", url)
	body, urls, err := fetcher.Fetch(url)

	//and update the status in a synced zone.
	fetched.Lock()
	fetched.m[url] = err
	fetched.Unlock()

	if err != nil {
		fmt.Printf("Error on %v: %v\n", url, err)
		return
	}
	fmt.Printf("Found: %s %q\n", url, body)

	done := make(chan bool)
	for i, u := range urls {
		fmt.Printf("-> Crawling child %v/%v of %v: %v.\n", i+1, len(urls), url, u)
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			done <- true
		}(u)
	}

	for i, u := range urls {
		fmt.Printf("<- [%v] %v/%v Waiting for child %v.\n", url, i+1, len(urls), u)
		<-done // block waiting
	}
	fmt.Printf("Done with %v\n", url)
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
	fmt.Println("Fetching stats\n--------------------")
	for url, err := range fetched.m {
		if err != nil {
			fmt.Printf("%v failed: %v\n", url, err)
		} else {
			fmt.Printf("%v was fetched!\n", url)
		}
	}
}

// fakeFetcher 是返回若干结果的 Fetcher。
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

// fetcher 是填充后的 fakeFetcher。
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


/* Console output:

Fetch: https://golang.org/
Found: https://golang.org/ "The Go Programming Language"
-> Crawling child 1/2 of https://golang.org/: https://golang.org/pkg/.
-> Crawling child 2/2 of https://golang.org/: https://golang.org/cmd/.
<- [https://golang.org/] 1/2 Waiting for child https://golang.org/pkg/.
Fetch: https://golang.org/cmd/
Error on https://golang.org/cmd/: not found: https://golang.org/cmd/
<- [https://golang.org/] 2/2 Waiting for child https://golang.org/cmd/.
Fetch: https://golang.org/pkg/
Found: https://golang.org/pkg/ "Packages"
-> Crawling child 1/4 of https://golang.org/pkg/: https://golang.org/.
-> Crawling child 2/4 of https://golang.org/pkg/: https://golang.org/cmd/.
-> Crawling child 3/4 of https://golang.org/pkg/: https://golang.org/pkg/fmt/.
-> Crawling child 4/4 of https://golang.org/pkg/: https://golang.org/pkg/os/.
<- [https://golang.org/pkg/] 1/4 Waiting for child https://golang.org/.
Fetch: https://golang.org/pkg/os/
Found: https://golang.org/pkg/os/ "Package os"
-> Crawling child 1/2 of https://golang.org/pkg/os/: https://golang.org/.
-> Crawling child 2/2 of https://golang.org/pkg/os/: https://golang.org/pkg/.
<- [https://golang.org/pkg/os/] 1/2 Waiting for child https://golang.org/.
Already fetched: https://golang.org/pkg/
<- [https://golang.org/pkg/os/] 2/2 Waiting for child https://golang.org/pkg/.
Already fetched: https://golang.org/
<- [https://golang.org/pkg/] 2/4 Waiting for child https://golang.org/cmd/.
Already fetched: https://golang.org/cmd/
<- [https://golang.org/pkg/] 3/4 Waiting for child https://golang.org/pkg/fmt/.
Fetch: https://golang.org/pkg/fmt/
Found: https://golang.org/pkg/fmt/ "Package fmt"
-> Crawling child 1/2 of https://golang.org/pkg/fmt/: https://golang.org/.
-> Crawling child 2/2 of https://golang.org/pkg/fmt/: https://golang.org/pkg/.
<- [https://golang.org/pkg/fmt/] 1/2 Waiting for child https://golang.org/.
Already fetched: https://golang.org/pkg/
<- [https://golang.org/pkg/fmt/] 2/2 Waiting for child https://golang.org/pkg/.
Already fetched: https://golang.org/
Done with https://golang.org/pkg/os/
<- [https://golang.org/pkg/] 4/4 Waiting for child https://golang.org/pkg/os/.
Already fetched: https://golang.org/
Done with https://golang.org/pkg/fmt/
Done with https://golang.org/pkg/
Done with https://golang.org/
Fetching stats
--------------------
https://golang.org/ was fetched!
https://golang.org/cmd/ failed: not found: https://golang.org/cmd/
https://golang.org/pkg/ was fetched!
https://golang.org/pkg/os/ was fetched!
https://golang.org/pkg/fmt/ was fetched!
*/
