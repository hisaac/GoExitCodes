package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	exitCode := run()
	fmt.Printf("Process finished with the exit code %d (%s)", exitCode.Int(), exitCode.String())
	os.Exit(exitCode.Int())
}

func run() ExitCode {
	// Return a random exit code
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		return Success
	} else {
		return Fail
	}
}
