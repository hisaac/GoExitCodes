//go:generate stringer -type=ExitCode

package main

type ExitCode int

const (
	Success ExitCode = iota
	Fail
)

func (i ExitCode) Int() int {
	return int(i)
}
