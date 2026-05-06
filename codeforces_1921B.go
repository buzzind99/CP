//go:build codeforces_1921B

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
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		s, f := next(), next()
		countA, countB := 0, 0
		for i := range n {
			if s[i] == '1' && f[i] == '0' {
				countA++
			} else if s[i] == '0' && f[i] == '1' {
				countB++
			}
		}

		fmt.Fprintln(wr, max(countA, countB))
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1921/B
  Tags: greedy, implementation
  Rating: 800
*/
