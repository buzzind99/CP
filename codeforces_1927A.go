//go:build codeforces_1927A

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
		s := next()
		first, last := -1, -1
		for i := range n {
			if s[i] == 'B' {
				if first == -1 { first = i }
				last = i
			}
		}

		fmt.Fprintln(wr, last-first+1)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1927/A
  Tags: greedy, strings
  Rating: 800
*/
