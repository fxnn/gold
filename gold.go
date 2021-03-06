package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fxnn/gold/import/csv"
	"github.com/fxnn/gold/model"
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

	var file *os.File
	var err error
	if file, err = os.Open(*source); err != nil {
		fmt.Printf("Error: couldn't open source: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var transactions []model.Transaction
	if transactions, err = csv.ReadSparkasse(file); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	for idx, t := range transactions {
		fmt.Printf("%3d %s: %d %s\n", idx, t.Partner, t.Amount, t.Currency)
	}

}
