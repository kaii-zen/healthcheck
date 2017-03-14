package main

import "fmt"

// Result holds a single healthcheck result. This includes a user-displaybale output as well as numeric values representing the
// result itself and the lower boundaries over which the result is considered either critical or a warning.
type Result struct {
	output            string
	value, warn, crit float64
}

// Value sets the value of a result, returning a pointer to said result.
func (r *Result) Value(v float64) *Result {
	r.value = v
	return r
}

// Crit sets the lower boundary over which a result value would be considered critical. returns a pointer to said result.
// If set to 0 (default), the result will never produce a CRIT status.
func (r *Result) Crit(c float64) *Result {
	r.crit = c
	return r
}

// Warn sets the lower boundary over which a result value would cause a warning. returns a pointer to said result.
// If set to 0 (default), the result will never produce a WARN status.
func (r *Result) Warn(w float64) *Result {
	r.warn = w
	return r
}

// Output sets the user-displayable output of the result. This is usually a formatted version of the result value,
// alongside any other data that might be relevant.
func (r *Result) Output(o string) *Result {
	r.output = o
	return r
}

// Status checks where value is in relation to crit and warn and returns a Status value accordingly.
func (r *Result) Status() Status {
	switch {
	case r.crit != 0 && r.value > r.crit:
		return CRIT
	case r.warn != 0 && r.value > r.warn:
		return WARN
	default:
		return OK
	}
}

// String returns a string representation of the result.
// This consists of the status in parentheses and the output. For example: "(CRIT) /var: 99%"
func (r Result) String() string {
	return fmt.Sprintf("(%v) %v", r.Status(), r.output)
}
