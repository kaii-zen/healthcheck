package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestResultsString(t *testing.T) {
	cases := []struct {
		results        *Results
		expectedString string
	}{
		{
			results: (new(Results).
				Add(new(Result).Value(75).Warn(80).Crit(90).Output("/: 75%")).
				Add(new(Result).Value(85).Warn(80).Crit(90).Output("/var: 85%")).
				Add(new(Result).Value(95).Warn(80).Crit(90).Output("/opt: 95%"))),
			expectedString: fmt.Sprintf("CRIT | (WARN) /var: 85%% | (CRIT) /opt: 95%%"),
		},
	}

	for _, c := range cases {
		if strings.Compare(c.results.String(), c.expectedString) != 0 {
			t.Errorf("Expected: %s; Actual: %s", c.expectedString, c.results)
		}
	}
}
