package main

type Status uint

const (
	OK = iota
	WARN
	CRIT
)

var Statuses = [...]string{
	"OK",
	"WARN",
	"CRIT",
}

func (self Status) String() string { return Statuses[self] }
