package main

// Status type represents that status of a single healthcheck result.
// That is, either OK, WARN or CRit
type Status uint

const (
	// OK represents a passing healthcheck
	OK = iota
	// WARN represents a the warning state of a healthcheck
	WARN
	// CRIT represents a the critical state of a healthcheck
	CRIT
)

// Statuses is a slice holding the string representations of Status
var Statuses = [...]string{
	"OK",
	"WARN",
	"CRIT",
}

// String returns a string representation of Status
func (s Status) String() string { return Statuses[s] }
