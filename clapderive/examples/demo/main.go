package main

import "fmt"

// Simple program to greet a person
//
//go:generate go run github.com/jcbhmr/go-clap/clapderive/Parser
type Args struct {
	// Name of the person to greet
	Name string `arg:"short,long"`
	// Number of times to greet
	Count uint8 `arg:"short,long,defaultValueT=1"`
}

// Check out Args-Parser.go to see the generated code!

func main() {
	args := ArgsParse()

	for i := uint8(0); i < args.Count; i++ {
		fmt.Printf("Hello, %s!\n", args.Name)
	}
}
