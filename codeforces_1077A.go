//go:build codeforces_1077A

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

func isEven(n int64) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt64()
	for range t {
		a, b, k := nextInt64(), nextInt64(), nextInt64()

		ans := (a - b) * (k / 2)
		if !isEven(k) { ans += a }

		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1077/A
  Tags: imlementation, math
  Rating: 800
*/
