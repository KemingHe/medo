package main

import (
	"fmt"
	"net/http"
	"os"
)

// main makes http request to google.com and prints response to console.
func main() {
	site := "https://www.example.com"
	res, getErr := http.Get(site)
	if getErr != nil {
		fmt.Printf("Error making GET request to %v: %v\n", site, getErr)
		os.Exit(1)
	}

	resReservedSize := 99999
	resBodySlice := make([]byte, resReservedSize)
	fmt.Printf("resBodySlice length = %v\n", len(resBodySlice))

	actualSize, readErr := res.Body.Read(resBodySlice)
	if readErr != nil {
		fmt.Printf("Error reading response body: %v\n", readErr)
		os.Exit(1)
	}

	fmt.Printf("Reserved %v, used %v\n", resReservedSize, actualSize)
	fmt.Println(string(resBodySlice))
}
