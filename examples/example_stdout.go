package main

import (
	"os"

	"github.com/he4d/simplejack"
)

func main() {
	logger := simplejack.New(simplejack.WARNING, os.Stdout)
	logger.Trace.Print("This will land in nirvana")
	logger.Warning.Print("This will print out to stdout")
	//Result:
	//WARNING: 2017/07/01 09:29:44 example_stdout.go:12: This will print out to stdout
}
