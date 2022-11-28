package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	exitCode := run()
	fmt.Printf("Process finished with the exit code %d (%s)", int(exitCode), exitCode.String())
	os.Exit(int(exitCode))
}

func run() ExitCode {
	// Return a random exit code
	rand.Seed(time.Now().UnixNano())
	return ExitCode(rand.Intn(2))
}
