//go:build codeforces_2000B

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
		first := nextInt()
		l, r := first, first
		followed := true
		for range n-1 {
			curr := nextInt()
			if curr == l-1 || curr == r+1 {
				l = min(l, curr)
				r = max(r, curr)
			} else {
				followed = false
			}
		}

		if followed {
			fmt.Fprintln(wr, "YES")
		} else {
			fmt.Fprintln(wr, "NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2000/B
  Tags: two pointers
  Rating: 800
*/
