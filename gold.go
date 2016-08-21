package main

import (
	"flag"
	"fmt"
	"os"
)

var source = flag.String("source", "", "Path to source CSV file")

func main() {

	flag.Parse()

	fmt.Printf("gold\n")
	fmt.Printf("source: %s\n", *source)
	fmt.Println()

	if *source == "" {
		fmt.Printf("Error: No source given\n")
		os.Exit(1)
	}

}
