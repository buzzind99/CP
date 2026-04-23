//go:build codeforces_1843C

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

func nextInt64() int64 {
	i, _ := strconv.ParseInt(next(), 10, 64)
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt64()
	for range t {
		n := nextInt64()
		sum := n
		for n > 0 {
			sum += n/2
			n /= 2
		}

		fmt.Fprintln(wr, sum)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1843/C
  Tags: bitmasks, combinatorics, math, trees
  Rating: 800
*/
