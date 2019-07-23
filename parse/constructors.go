package parse

import "github.com/timraymond/timtoml/ledger"

// newTXs creates a new sequence of transactions
func newTXs(ents interface{}) ([]ledger.TX, error) {
	out := []ledger.TX{}
	entries, err := toII(ents)
	if err != nil {
		return out, err
	}
	for _, ent := range entries {
		tx, err := toTX(ent)
		if err != nil {
			return out, err
		}
		out = append(out, tx)
	}
	return out, nil
}

// newTX creates a new transaction
func newTX(date, state, payee, postings interface{}) (ledger.TX, error) {
	dt, err := toTime(date)
	if err != nil {
		return ledger.TX{}, nil
	}

	py, err := toString(payee)
	if err != nil {
		return ledger.TX{}, nil
	}

	return ledger.TX{
		Date:  dt,
		Payee: py,
	}, nil
}
