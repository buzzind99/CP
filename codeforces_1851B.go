//go:build codeforces_1851B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
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
		orig := make([]int, n)
		sorted := make([]int, n)
		for i := range n {
			val := nextInt()
			orig[i] = val
			sorted[i] = val
		}
		slices.Sort(sorted)

		possible := true
		for i := range n {
			if orig[i]%2 != sorted[i]%2 {
				possible = false
				break
			}
		}

		if possible {
			fmt.Fprintln(wr, "YES")
		} else {
			fmt.Fprintln(wr, "NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1851/B
  Tags: greedy, sortings, two pointers
  Rating: 800
*/
