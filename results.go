package main

import (
	"strings"
)

type Results []*Result

func (results *Results) Add(result *Result) *Results {
	*results = append(*results, result)
	return results
}

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
