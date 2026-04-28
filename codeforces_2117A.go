//go:build codeforces_2117A

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
		n, x := nextInt(), nextInt()
		doors := make([]int, n)
		for i := range n { doors[i] = nextInt() }

		firstClosed := -1
		for i, d := range doors {
			if d == 1 { firstClosed = i; break }
		}

		if firstClosed == -1 {
			fmt.Fprintln(wr, "YES")
			continue
		}

		possible := true
		for i := firstClosed + x; i < n; i++ {
			if doors[i] == 1 { possible = false; break }
		}
		if possible { fmt.Fprintln(wr, "YES") } else { fmt.Fprintln(wr, "NO") }
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2117/A
  Tags: greedy, implementation
  Rating: 800
*/
