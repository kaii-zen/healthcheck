package main

import (
	"testing"
)

func TestResultStatus(t *testing.T) {
	cases := []struct {
		result         *Result
		expectedStatus Status
	}{
		{
			result:         new(Result).Value(75).Warn(80).Crit(90),
			expectedStatus: OK,
		},
		{
			result:         new(Result).Value(85).Warn(80).Crit(90),
			expectedStatus: WARN,
		},
		{
			result:         new(Result).Value(95).Warn(80).Crit(90),
			expectedStatus: CRIT,
		},
		{
			result:         new(Result).Value(85).Warn(80),
			expectedStatus: WARN,
		},
	}

	for _, c := range cases {
		if c.result.Status() != c.expectedStatus {
			t.Errorf("Expected: %s; Actual: %s", c.expectedStatus, c.result.Status())
		}
	}
}
