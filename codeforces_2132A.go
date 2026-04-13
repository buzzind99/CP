//go:build codeforces_2132A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		_, a := nextInt(), next()
		m, b := nextInt(), next()
		c := next()
		sol := a
		for i := range m {
			if c[i] == 'D' {
				sol += string(b[i])
			} else {
				sol = string(b[i])+sol
			}
		}

		fmt.Fprintln(wr, sol)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2132/A
  Tags: brute force, implementation, strings
  Rating: 800
*/
