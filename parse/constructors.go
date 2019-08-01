package parse

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/timraymond/timtoml/ledger"
)

// newTXs creates a new sequence of transactions
func newTXs(ents interface{}) ([]ledger.TX, error) {
	out := []ledger.TX{}
	entries, err := toII(ents)
	if err != nil {
		return out, errors.Wrap(err, "toII")
	}
	for _, ent := range entries {
		tx, err := toTX(ent)
		if err != nil {
			return out, errors.Wrap(err, "toTX")
		}
		out = append(out, tx)
	}
	return out, nil
}

// newTX creates a new transaction
func newTX(date, state, payee, postings interface{}) (ledger.TX, error) {
	dt, err := toTime(date)
	if err != nil {
		return ledger.TX{}, errors.Wrap(err, "toTime")
	}

	py, err := toString(payee)
	if err != nil {
		return ledger.TX{}, errors.Wrap(err, "toString")
	}

	posts, err := toPostings(postings)
	if err != nil {
		return ledger.TX{}, errors.Wrap(err, "toPostings")
	}

	return ledger.TX{
		Date:     dt,
		Payee:    py,
		Postings: posts,
	}, nil
}

// newPost creates a new posting
func newPost(st, act, amt interface{}) (ledger.Posting, error) {
	state := toState(st)

	account, err := toString(act)
	if err != nil {
		return ledger.Posting{}, errors.Wrap(err, "toString")
	}

	amount, err := toAmount(amt)
	if err != nil {
		return ledger.Posting{}, errors.Wrap(err, "toAmount")
	}

	return ledger.Posting{
		State:    state,
		Account:  account,
		Amount:   amount,
		Currency: "USD",
	}, nil
}

func newState(sig interface{}) (ledger.PostingState, error) {
	return ledger.StateUncleared, nil
}

func newAmount(neg, un, cts interface{}) (int, error) {
	_, negative := neg.([]byte)

	total := 0

	uStr, err := toByteSlice(un)
	if err != nil {
		return 0, errors.Wrap(err, "toByteSlice units")
	}

	units, err := strconv.Atoi(string(uStr))
	if err != nil {
		return 0, errors.Wrap(err, "strconv.Atoi units")
	}

	total += units

	// this is to accommodate cents
	total = total * 100

	cStr, err := toByteSlice(cts)
	if err != nil {
		return 0, errors.Wrap(err, "toByteSliceFromII cents")
	}

	cents, err := strconv.Atoi(string(cStr[1:])) // remember the leading "."
	if err != nil {
		return 0, errors.Wrap(err, "strconv.Atoi cents")
	}

	total += cents

	if negative {
		total = total * -1
	}
	return total, nil
}
