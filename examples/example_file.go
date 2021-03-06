package main

import (
	"log"
	"os"

	"github.com/he4d/simplejack"
)

func main() {
	file, err := os.OpenFile("example.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file example.log : %s", err)
	}

	logger := simplejack.New(simplejack.DEBUG, file)
	logger.Trace.Print("This will land in nirvana")
	logger.Fatal.Fatal("This will write to the file example.log and os.Exit(1)")
}
