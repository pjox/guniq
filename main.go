package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cespare/xxhash"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage:\n\t%s [input] [output]\n", os.Args[0])
		os.Exit(1)
	}

	in, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	out, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
	bufin := bufio.NewReader(in)
	bufout := bufio.NewWriter(out)

	tab := make(map[uint64]int)

	for par, err := bufin.ReadString('\n'); err == nil; par, err = bufin.ReadString('\n') {
		hash := xxhash.Sum64String(par)
		if _, ok := tab[hash]; !ok || par == "\n" {
			tab[hash] = 1
			bufout.WriteString(par)
		}
	}
	bufout.Flush()
	in.Close()
	out.Close()
}
