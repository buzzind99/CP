//go:build codeforces_2149A

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
		n := nextInt()
		sol, negatives := 0, 0
		for range n {
			a := nextInt()
			if a == 0 { sol++ } else if a < 0 { negatives++ }
		}
		sol += 2*(negatives%2)

		fmt.Fprintln(wr, sol)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2149/A
  Tags: math
  Rating: 800
*/
