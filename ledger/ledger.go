package ledger

import "time"

// TX represents a ledger transaction
type TX struct {
	Date     time.Time
	Payee    string
	Postings []Posting
}

// Posting represents a transfer into or out of an account
type Posting struct {
	Account  string
	Amount   int
	Currency string
	State    PostingState
}

type PostingState int

const (
	StateInvalid PostingState = iota
	StateUncleared
	StateCleared
	StatePending
)
