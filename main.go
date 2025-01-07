package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/steelx/extractlinks"
)

var (
	config = &tls.Config{
		InsecureSkipVerify: true,
	}

	transport = &http.Transport{
		TLSClientConfig: config,
	}

	netClient = &http.Client{
		Transport: transport,
	}
	queue     = make(chan string, 100)
	hasVisted = make(map[string]bool)
	mu        sync.Mutex
)

func main() {

	eCommerce := []string{
		"https://www.hostinger.com/",
		"https://www.wix.com/",
		"https://www.bigcartel.com/",
		"https://www.etsy.com/",
		"https://www.squarespace.com/",
		"https://www.weebly.com/",
		"https://www.shopify.com/",
		"https://www.volusion.com/",
		"https://www.bigcommerce.com/",
		"https://www.ecwid.com/",
	}
	const numWorkers = 12
	var wg sync.WaitGroup
	//spawning worker goroutines

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	for _, webSiteUrl := range eCommerce {
		queue <- webSiteUrl
	}

	wg.Wait()
	close(queue)
	fmt.Println("Crawling completed!")

}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range queue {
		mu.Lock()
		if !hasVisted[url] {
			hasVisted[url] = true
			mu.Unlock()
			ProductUrls(url)
		} else {
			mu.Unlock()
		}
	}
}

func ProductUrls(href string) {
	hasVisted[href] = true
	fmt.Printf("products list  Urls --> %v \n", href)

	resp, err := netClient.Get(href)
	checkErr(err)
	if resp == nil {
		fmt.Printf("Received nil response for URL: %s\n", href)
		return
	}
	defer resp.Body.Close()

	links, err := extractlinks.All(resp.Body)
	checkErr(err)

	for _, link := range links {

		absoluteUrl := MakeCompleteProductUrl(link.Href, href)
		if absoluteUrl == "" || !isValidURL(absoluteUrl) {
			continue
		}

		go func() {
			queue <- absoluteUrl
		}()
	}
}

func MakeCompleteProductUrl(href, baseUrl string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	base, err := url.Parse(baseUrl)
	if err != nil {
		return ""
	}
	toFixedUri := base.ResolveReference(uri)
	return toFixedUri.String()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Skipping unsupported URL or error occurred:", err)
		return
	}
}

func isValidURL(href string) bool {
	parsedURL, err := url.Parse(href)
	return err == nil && (parsedURL.Scheme == "http" || parsedURL.Scheme == "https")
}
