package main

import (
	"flag"
	"fmt"
)

var mg = flag.String("must-gather", "must-gather.tar", "The path to the must-father.tar file to analyze")

func init() {
	flag.StringVar(mg, "m", "must-gather.tar", "The path to the must-father.tar file to analyze")
}

func main() {
	flag.Parse()
	fmt.Printf("Must-gather: %s\n", *mg)
}
