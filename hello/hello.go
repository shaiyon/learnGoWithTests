package main

import "fmt"

func helloPrefixes() map[string]string {
    return map[string]string{
        "English": "Hello, ",
        "Spanish": "Hola, ",
        "French":  "Bonjour, ",
    }
}

type HelloOptions struct {
	name string
	language string
}

func Hello(options HelloOptions) string {
	if options.name == "" {
		options.name = "World"
	}
	if options.language == "" {
		options.language = "English"
	}

	prefixes := helloPrefixes()
	return prefixes[options.language] + options.name
}

func main() {
	fmt.Println(Hello(HelloOptions{name: "Chris"}))
}