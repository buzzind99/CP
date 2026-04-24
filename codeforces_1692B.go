//go:build codeforces_1692B

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

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		m := make(map[int]bool)
		dupes := 0
		for range n {
			a := nextInt()
			if !m[a] { m[a] = true } else { dupes++ }
		}

		if isEven(dupes) {
			fmt.Fprintln(wr, n-dupes)
		} else {
			fmt.Fprintln(wr, n-dupes-1)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1692/B
  Tags: greedy, sortings
  Rating: 800
*/
