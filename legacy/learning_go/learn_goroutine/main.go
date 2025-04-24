package main

import (
	"fmt"
	"net/http"
	// "sync"
	"time"
)

// Note: concurrency (multi-threading) is not parallelism (actually utilizing multiple cores)

func main() {
	sites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	// Solution 1: goroutine and WaitGroup (channel-less) simple implementation
	// var wg sync.WaitGroup
	// for _, site := range sites {
	// 	wg.Add(1) // Increment counter at start of every new goroutine
	// 	go func(site string) {
	// 		defer wg.Done() // Make sure Done is called even if gorountine panics
	// 		status := getSiteStatus(site)
	// 		reportSiteStatus(site, status)
	// 	}(site)
	// }
	// wg.Wait() // Wait for all goroutines to finish before main exits

	// // Solution 2: goroutine with simple channel
	// // (channels are the only way to communicate between goroutines; plus, channels are typed)
	// done := make(chan bool)
	// for _, site := range sites {
	// 	go func (site string) {
	// 		status := getSiteStatus(site)
	// 		reportSiteStatus(site, status)
	// 		done <- true
	// 	}(site)
	// }

	// // Main goroutine will use done channel to wait for all child goroutines to finish
	// for range len(sites) { // Use range int for modernized for loop in go
	// 	<-done
	// }

	// Solution 3: repeating goroutines using channels to restart once finished
	c := make(chan string)
	for _, site := range sites {
		go func (site string) {
			status := getSiteStatus(site)
			reportSiteStatus(site, status)
			c <- site
		}(site)
	}

	// // Efficient but confusing, don't do this
	// for { // Use empty for for infinite loop
	// 	go func (site string) {
	// 		status := getSiteStatus(site)
	// 		reportSiteStatus(site, status)
	// 		c <- site
	// 	}(<-c)
	// }

	// Use this instead, "range c" waits for messages in channel continuously
	for site := range c {
		go func (site string) {
			time.Sleep(5 * time.Second) // For each new goroutine after first batch, add 5s delay to avoid overwhelming server
			status := getSiteStatus(site)
			reportSiteStatus(site, status)
			c <- site
		}(site)
	}
}

// getSiteStatus returns "up" if site is up, "down" if site is down, and "error" if error accessing site
func getSiteStatus(site string) string {
	res, err := http.Get(site)
	if err != nil {
		fmt.Printf("Error checking %s: %v\n", site, err)
		return "error"
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "down"
	}

	return "up"
}

// reportSiteStatus prints the site and corresponding status to stdout
func reportSiteStatus(site string, status string) {
	fmt.Printf("Site %q has status %q\n", site, status)
}
