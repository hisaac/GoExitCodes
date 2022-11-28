//go:generate stringer -type=ExitCode

package main

type ExitCode int

const (
	Success ExitCode = 0
	Fail    ExitCode = 1
)
