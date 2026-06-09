//go:build codeforces_1772B

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

func isBeautiful(a, b, c, d int) bool {
	return a < b && c < d && a < c && b < d
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		a, b := nextInt(), nextInt()
		c, d := nextInt(), nextInt()

		possible := false
		for range 4 {
			if isBeautiful(a, b, c, d) {
				possible = true
				break
			}
			a, b, d, c = c, a, b, d
		}

		if possible {
			fmt.Fprintln(wr, "YES")
		} else {
			fmt.Fprintln(wr, "NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1772/B
  Tags: brute force, implementation
  Rating: 800
*/
