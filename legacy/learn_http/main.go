package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// main makes http request to google.com and prints response body to console.
func main() {
	site := "https://www.example.com"
	res, getErr := http.Get(site)
	if getErr != nil {
		fmt.Printf("Error making GET request to %v: %v\n", site, getErr)
		os.Exit(1)
	}
	// Defer ensures Body.Close() executes when the function returns, regardless of how it exits
	// This prevents resource leaks (file descriptors) and allows TCP connection reuse
	// Placing defer immediately after response validation ensures cleanup occurs on all code paths
	defer res.Body.Close()

	// Use the more efficient io.ReadAll instead of res.Body.Read to avoid fix buffer and EOF issues
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Printf("Error reading response body: %v\n", readErr)
		os.Exit(1)
	}

	fmt.Printf("Response size: %v bytes\n", len(body))
	fmt.Println(string(body))
}
