package main

import (
	"archive/tar"
	"bufio"
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

	r := bufio.NewReader(f)
	f2b, err := r.Peek(2) //read 2 bytes
	if f2b[0] == 31 && f2b[1] == 139 {
		log.Fatalf("File %s is gzipped", *mg)
	}

	f.Seek(0, 0)
	tr := tar.NewReader(f)
	h, err := tr.Next()
	if err != nil {
		log.Fatalf("Could not untar file: %s", err)
	}

	fmt.Printf("Tar reader: %v", tr)
	fmt.Printf("Tar header: %v", h)

}
