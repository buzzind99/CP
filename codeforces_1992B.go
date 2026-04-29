//go:build codeforces_1992B

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
		_, k := nextInt64(), nextInt64()
		var maxVal, sum int64 = -1, 0
		for range k {
			a := nextInt64()
			sum += 2*a-1
			maxVal = max(maxVal, a)
		}
		ans := sum-2*maxVal+1

		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1992/B
  Tags: greedy, math, sortings
  Rating: 800
*/
