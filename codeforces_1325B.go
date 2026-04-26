//go:build codeforces_1325B

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
		m := make(map[int]struct{}, n)
		for range n {
			a := nextInt()
			m[a] = struct{}{}
		}

		fmt.Fprintln(wr, len(m))
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1325/B
  Tags: greedy, implementation
  Rating: 800
*/
