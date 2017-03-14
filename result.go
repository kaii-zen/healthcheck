package main

import "fmt"

type Result struct {
	output            string
	value, warn, crit float64
}

func (r *Result) Value(v float64) *Result {
	r.value = v
	return r
}

func (r *Result) Crit(c float64) *Result {
	r.crit = c
	return r
}

func (r *Result) Warn(w float64) *Result {
	r.warn = w
	return r
}

func (r *Result) Output(o string) *Result {
	r.output = o
	return r
}

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

func (r Result) String() string {
	return fmt.Sprintf("(%v) %v", r.Status(), r.output)
}
