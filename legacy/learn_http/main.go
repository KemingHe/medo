package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

// main makes an http request to a given site and prints response body to console.
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

	// Use the more efficient io.ReadAll instead of res.Body.Read to avoid fixed buffer and EOF issues
	// body, readErr := io.ReadAll(res.Body)
	// if readErr != nil {
	// 	fmt.Printf("Error reading response body: %v\n", readErr)
	// 	os.Exit(1)
	// }

	// fmt.Println(string(body))
	// fmt.Printf("Response size: %v bytes\n", len(body))

	// Alternatively, directly use the os.Stdout Writer to print res.Body Reader to console
	// copyCount, writeErr := io.Copy(os.Stdout, res.Body)
	// if writeErr != nil {
	// 	fmt.Printf("Error writing res body to console: %v\n", writeErr)
	// 	os.Exit(1)
	// }
	// fmt.Printf("Copied %v bytes from res.Body to console\n", copyCount)

	// Practice: using custom Writer
	lwPtr := &logWriter{}
	logLen, _ := io.Copy(lwPtr, res.Body) // Error will always be nil, ignored
	fmt.Printf("Printed %v bytes of data from response\n", logLen)
}

func (*logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}