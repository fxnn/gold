package model

import (
	"time"
)

// Transaction models a single payment done from or to the user's account.
type Transaction struct {

	// date of the transaction
	Date time.Time

	// postingText describes the contents of this transaction
	PostingText string

	// reference code allows to uniquely identify the reason of this transaction
	Reference string

	// parnter describes the other party (creditor/debitor) for this transaction
	Partner string

	// amount of minor units this transaction contains
	Amount int

	// currency of the amount in ISO three letter notation
	Currency string
}
