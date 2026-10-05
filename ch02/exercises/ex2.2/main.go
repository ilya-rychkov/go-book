package main

import (
	"bufio"
	"ex2.2/unitconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := unitconv.Foot(t)
			m := unitconv.Metre(t)
			p := unitconv.Pound(t)
			k := unitconv.Kilogram(t)
			fmt.Printf("%s = %s, %s = %s,\n%s = %s, %s = %s\n",
				f, unitconv.FToM(f), m, unitconv.MtoF(m), p, unitconv.PToK(p), k, unitconv.KToP(k))
		}
	} else {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := unitconv.Foot(t)
			m := unitconv.Metre(t)
			p := unitconv.Pound(t)
			k := unitconv.Kilogram(t)
			fmt.Printf("%s = %s, %s = %s,\n%s = %s, %s = %s\n",
				f, unitconv.FToM(f), m, unitconv.MtoF(m), p, unitconv.PToK(p), k, unitconv.KToP(k))
		}
	}
}
