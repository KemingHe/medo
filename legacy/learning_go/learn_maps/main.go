package main

import "fmt"

// Maps are of reference type in go, no need to explicitly pass by pointer in functions
// All keys of a map must be the same type
// All values of a map must be the same type
// Keys and values can be different types

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	// Adding a new key-value pair to a map
	// Must use square brace syntax here because all key-values are typed in go
	colors["white"] = "#ffffff"

	// Deleting an existing key-value pair from a map
	delete(colors, "white")

	printColorMap(colors)
}

func printColorMap(c map[string]string) {
	// Keys in a map are indexed and can be iterated over
	for color, hex := range c {
		fmt.Printf("Color %v has hex code: %v\n", color, hex)
	}
}