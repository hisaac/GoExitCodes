//go:generate stringer -type=ExitCode

package main

type ExitCode int

const (
	Success ExitCode = 0
	Failure ExitCode = 1
)
