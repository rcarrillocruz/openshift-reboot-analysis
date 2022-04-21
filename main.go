package main

import (
	"archive/tar"
	"flag"
	"fmt"
	"log"
	"os"
)

var mg = flag.String("must-gather", "must-gather.tar", "The path to the must-father.tar file to analyze")

func init() {
	flag.StringVar(mg, "m", "must-gather.tar", "The path to the must-father.tar file to analyze")
}

func main() {
	flag.Parse()
	fmt.Printf("Must-gather: %s\n", *mg)

	f, err := os.Open(*mg)
	if err != nil {
		log.Fatal("Could not open file")
	}

	tr := tar.NewReader(f)
	h, err := tr.Next()
	if err != nil {
		log.Fatalf("Could not untar file: %s", err)
	}

	fmt.Printf("Tar reader: %v", tr)
	fmt.Printf("Tar header: %v", h)

}
