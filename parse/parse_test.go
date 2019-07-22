package parse_test

import (
	"testing"

	"github.com/timraymond/timtoml/parse"
)

func TestParse(t *testing.T) {
	parseTests := []struct {
		name      string
		config    string
		shouldErr bool
	}{
		{
			"empty",
			"",
			false,
		},
	}

	for _, test := range parseTests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := parse.Parse(test.name, []byte(test.config))
			if err != nil && !test.shouldErr {
				t.Fatal("Unexpected err:", err)
			}

			if err == nil && test.shouldErr {
				t.Fatal("Expected an err but received none")
			}
		})
	}
}
