package main

import (
	"time"
)

type transaction struct {

	// date of the transaction
	date time.Time

	// postingText describes the contents of this transaction
	postingText string

	// reference code allows to uniquely identify the reason of this transaction
	reference string

	// parnter describes the other party (creditor/debitor) for this transaction
	partner string

	// amount of minor units this transaction contains
	amount int

	// currency of the amount in ISO three letter notation
	currency string
}
