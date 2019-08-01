package parse

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/timraymond/ledger/ast"
)

func toII(ii interface{}) ([]interface{}, error) {
	out, ok := ii.([]interface{})
	if !ok {
		return out, fmt.Errorf("Expected []interface{}, but got %T", ii)
	}
	return out, nil
}

func toTX(tx interface{}) (ast.TX, error) {
	out, ok := tx.(ast.TX)
	if !ok {
		return out, fmt.Errorf("Expected ast.TX, but got %T", tx)
	}
	return out, nil
}

func toTime(tm interface{}) (time.Time, error) {
	out, ok := tm.(time.Time)
	if !ok {
		return out, fmt.Errorf("Expected time.Time, but got %T", tm)
	}
	return out, nil
}

func toString(str interface{}) (string, error) {
	out, ok := str.(string)
	if !ok {
		return out, fmt.Errorf("Expected string, but got %T", str)
	}
	return out, nil
}

func toState(st interface{}) ast.PostingState {
	state, ok := st.(ast.PostingState)
	if !ok {
		return ast.StateUncleared
	}
	return state
}

func toPostings(psts interface{}) ([]ast.Posting, error) {
	ps, err := toII(psts)
	if err != nil {
		return []ast.Posting{}, errors.Wrap(err, "toII")
	}
	out := make([]ast.Posting, 0, len(ps))
	for _, p := range ps {
		post, ok := p.(ast.Posting)
		if !ok {
			return out, fmt.Errorf("Expected ast.Posting, but got %T", psts)
		}
		out = append(out, post)
	}
	return out, nil
}

func toByteSlice(sl interface{}) ([]byte, error) {
	bslice, ok := sl.([]byte)
	if !ok {
		return []byte{}, fmt.Errorf("Expected []byte, but got %T", sl)
	}
	return []byte(bslice), nil
}

func toAmount(amt interface{}) (int, error) {
	amount, ok := amt.(int)
	if !ok {
		return 0, fmt.Errorf("Expected int, but got %T", amt)
	}
	return amount, nil
}
