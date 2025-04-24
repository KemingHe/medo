package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	ebPTR := &englishBot{}
	sbPTR := &spanishBot{}

	printGreeting(ebPTR)
	printGreeting(sbPTR)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// In general, use pointer receivers to avoid large copies and future extendability for mutators

// Omit receiver instance if unused
func (*englishBot) getGreeting() string {
	return "hello"
}

func (*spanishBot) getGreeting() string {
	return "hola"
}
