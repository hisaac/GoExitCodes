package main

type ExitCode int

const (
	Success ExitCode = iota
	Fail
)

func (e ExitCode) Int() int {
	return int(e)
}

func (e ExitCode) String() string {
	switch e {
	case Success:
		return "Success"
	case Fail:
		return "Fail"
	}
	return "Unknown"
}
