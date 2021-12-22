package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	printNumber := flag.Bool("n", false, "print line number option")
	flag.Parse()
	line := 0
	for _, v := range flag.Args() {
		f, err := os.Open(v)
		if err != nil {
			log.Fatal(err)
			return
		}
		s := bufio.NewScanner(f)
		for s.Scan() {
			line += 1
			if *printNumber {
				fmt.Println(line, ":", s.Text())
			} else {
				fmt.Println(s.Text())
			}
		}
		if s.Err() != nil {
			// non-EOF error.
			log.Fatal(s.Err())
		}
	}
}
