package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	// "path/filepath"
)

func readFile(fn string) ([]string, error) {
	var lines []string
	f, err := os.Open(fn)
	if err != nil {
		return lines, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	defer f.Close()
	return lines, err
}

var n *bool = flag.Bool("n", false, "If you want line number, set true.")

func main() {
	var lineNum int = 1
	flag.Parse()
	for _, fn := range flag.Args() {
		lines, err := readFile(fn)
		if err != nil {
			log.Fatal(err)
		}

		for _, l := range lines {
			if *n {
				fmt.Fprintln(os.Stdout, strconv.Itoa(lineNum), ": ", l)
			} else {
				fmt.Fprintln(os.Stdout, l)
			}
			lineNum++
		}
	}
}
