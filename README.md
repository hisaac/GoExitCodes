# GoExitCodes

A simple experiment for adding an idiomatic way of handling exit codes in Go.

## Background

Most of my development experience has been on Apple platforms, and this little repo was born from me missing some of Swift/ObjC's more idiomatic/expressive tendencies.

For example, when exiting from a CLI built in Swift, you would do something like:

```swfit
exit(EXIT_SUCCESS)
```

or 

```swift
exit(EXIT_FAILURE)
```

`EXIT_SUCCESS` and `EXIT_FAILURE` are just macros for `0` and `1`, respectively, so nothing fancy at all, but it's just a more expressive way to represent exit codes.

## Usage

At its simplest, you can use `GoExitCodes` like so:

```go
func main() {
	exitCode := ExitCode.Success
	fmt.Printf("Process finished with the exit code %d (%s)", exitCode.Int(), exitCode.String())
	os.Exit(exitCode.Int())
}
```
