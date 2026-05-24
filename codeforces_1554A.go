//go:build codeforces_1554A

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

func nextInt64() int64 {
	i, _ := strconv.ParseInt(next(), 10, 64)
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		var prev, maxAns int64 = 0, 0
		for i := 0; i < n; i++ {
			current := nextInt64()
			if i > 0 {
				product := prev*current
				if product > maxAns {
					maxAns = product
				}
			}
			prev = current
		}

		fmt.Fprintln(wr, maxAns)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1554/A
  Tags: greedy
  Rating: 800
*/
