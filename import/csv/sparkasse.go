package csv

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/fxnn/gold/model"
)

import "io"

import "time"

// ReadSparkasse parses the MT940-CSV format as provided on
// sparkasse-gera-greiz.de
func ReadSparkasse(source *os.File) ([]model.Transaction, error) {

	csvReader := csv.NewReader(source)
	csvReader.Comma = ';'

	var result []model.Transaction

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			return result, nil
		} else if err != nil {
			return nil, fmt.Errorf("could not read CSV file: %s\n", err)
		}

		var t model.Transaction
		t.Date, err = time.Parse("02.01.06", record[2])
		if err != nil {
			fmt.Printf("Warning: couldn't parse date field: %s\n", err)
		}
		t.PostingText = record[3]
		t.Reference = record[4]
		t.Partner = record[5]
		t.Amount, err = parseAmount(record[8])
		if err != nil {
			fmt.Printf("Warning: couldn't parse amount field: %s\n", err)
		}
		t.Currency = record[9]

		result = append(result, t)
	}
}

func parseAmount(s string) (int, error) {
	if s == "" {
		return 0, fmt.Errorf("value is empty")
	}

	var sign int
	if s[0] == '-' {
		sign = -1
	} else {
		sign = 1
	}

	var commaIdx = strings.Index(s, ",")
	if commaIdx < 0 {
		return 0, fmt.Errorf("no separator of minor units")
	}

	var majorStr = s[0:commaIdx]
	if len(majorStr) == 0 {
		return 0, fmt.Errorf("value part for major units is empty")
	}
	var minorStr = s[commaIdx+1:]
	if len(minorStr) == 0 {
		return 0, fmt.Errorf("value part for minor units is empty")
	}
	var majorMult = int(math.Pow10(len(minorStr)))

	major, err := strconv.Atoi(majorStr)
	if err != nil {
		return 0, fmt.Errorf("major units could not be parsed: %s", err)
	}
	minor, err := strconv.Atoi(minorStr)
	if err != nil {
		return 0, fmt.Errorf("minor units could not be parsed: %s", err)
	}

	return sign * (sign*major*majorMult + minor), nil
}
