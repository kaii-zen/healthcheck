package main

import (
	"strings"
)

// Results type represents a collection of healthcheck results. The idea is that a single healthcheck may have
// more than one result, for example in the case of the disk usage check, we might care about more than one partition.
// In this case, the most severe status should bubble up. So if one partition gets CRIT, the entire healthcheck is critical.
type Results []*Result

// Add adds a single result to the collection, returning a pointer to the collection.
func (results *Results) Add(result *Result) *Results {
	*results = append(*results, result)
	return results
}

// Status checks for the statuses of each results, returning the most severe.
func (results Results) Status() (status Status) {
	status = OK
	for _, result := range results {
		switch result.Status() {
		case CRIT:
			status = CRIT
			return
		case WARN:
			status = WARN
		}
	}
	return
}

// Returns a string representation of the results.
// This consists of the most severe status as well as outputs of all checks that are not OK.
func (results Results) String() string {
	const separator = " | "
	var resultStrings []string
	for _, result := range results {
		if result.Status() == OK {
			continue
		}
		resultStrings = append(resultStrings, result.String())
	}

	resultStrings = append([]string{results.Status().String()}, resultStrings...)
	return strings.Join(resultStrings, separator)
}
