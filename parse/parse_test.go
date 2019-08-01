package parse_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/timraymond/ledger/ast"
	"github.com/timraymond/ledger/parse"
)

func TestParse(t *testing.T) {
	parseTests := []struct {
		name      string
		config    string
		expTxns   int // the expected number of transactions
		shouldErr bool
	}{
		{
			"empty",
			"",
			0,
			false,
		},
		{
			"basic",
			`2012-03-10 KFC
    Expenses:Food                $20.00
    Assets:Cash                 $-20.00
`,
			1,
			false,
		},
		{
			"basic_fail",
			`2012-03-10 KFC
Expenses:Food                $20.00
Assets:Cash                 $-20.00
`,
			1,
			true,
		},
		{
			"multiple",
			`2012-03-10 KFC
    Expenses:Food                $20.00
    Assets:Cash                 $-20.00
2012-03-11 KFC
    Expenses:Food                $20.00
    Assets:Cash                 $-20.00
`,
			2,
			false,
		},
		{
			"header_state",
			`2012-03-10 * KFC
    Expenses:Food                $20.00
    Assets:Cash                 $-20.00
`,
			1,
			false,
		},
		{
			"posting_state",
			`2012-03-10  KFC
    * Expenses:Food                $20.00
    * Assets:Cash                 $-20.00
`,
			1,
			false,
		},
	}

	for _, test := range parseTests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rawTxns, err := parse.Parse(test.name, []byte(test.config))
			if err != nil && !test.shouldErr {
				t.Fatal("Unexpected err:", err)
			}

			if err == nil && test.shouldErr {
				t.Fatal("Expected an err but received none")
			}

			if err != nil && test.shouldErr {
				return
			}

			txns, ok := rawTxns.([]ast.TX)
			if !ok {
				t.Fatalf("Expected []ast.TX, but got %T", rawTxns)
			}

			if len(txns) != test.expTxns {
				t.Fatalf("Expected %d transactions but received %d", test.expTxns, len(txns))
			}
		})
	}
}

func TestParse_Details(t *testing.T) {
	exp := []ast.TX{
		{
			Date:  time.Date(2012, 3, 10, 0, 0, 0, 0, time.UTC),
			Payee: "KFC",
			Postings: []ast.Posting{
				{
					Account:  "Expenses:Food",
					Amount:   2000,
					Currency: "USD",
					State:    ast.StateUncleared,
				},
				{
					Account:  "Assets:Cash",
					Amount:   -2000,
					Currency: "USD",
					State:    ast.StateUncleared,
				},
			},
		},
	}

	got, err := parse.ParseFile("./testdata/simple.dat")
	if err != nil {
		t.Fatal("Err opening file: err:", err)
	}

	if !cmp.Equal(exp, got) {
		t.Error("Parsed file differs: diff:", cmp.Diff(exp, got))
	}
}
