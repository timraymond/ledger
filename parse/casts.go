package parse

import (
	"fmt"
	"time"

	"github.com/timraymond/timtoml/ledger"
)

func toII(ii interface{}) ([]interface{}, error) {
	out, ok := ii.([]interface{})
	if !ok {
		return out, fmt.Errorf("Expected []interface{}, but got %T", ii)
	}
	return out, nil
}

func toTX(tx interface{}) (ledger.TX, error) {
	out, ok := tx.(ledger.TX)
	if !ok {
		return out, fmt.Errorf("Expected ledger.TX, but got %T", tx)
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
