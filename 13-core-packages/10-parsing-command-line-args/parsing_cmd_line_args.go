package main

import (
	"flag"
	"log"
	"math/rand"
)

// run with: go run -v parsing_cmd_line_args.go -max=100
func main() {
	// define flags
	maxp := flag.Int("max", 6, "the max value")

	// parse
	flag.Parse()

	// generate a number between 0 and max
	log.Println(rand.Intn(*maxp))
}
