//go:build codeforces_1926B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func abs(x int) int {
	if x < 0 { return -x }
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		maxCount, minCount := -1, 11
		for range n {
			s := next()
			count := strings.Count(s, "1")
			if count > 0 {
				maxCount = max(maxCount, count)
				minCount = min(minCount, count)
			}
		}

		if maxCount == minCount {
			fmt.Fprintln(wr, "SQUARE")
		} else {
			fmt.Fprintln(wr, "TRIANGLE")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1926/B
  Tags: geometry, implementation
  Rating: 800
*/
