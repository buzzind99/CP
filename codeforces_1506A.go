//go:build codeforces_1506A

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
		n, m, x := nextInt64(), nextInt64(), nextInt64()
		x--
		row := x % n
		col := x / n
		ans := row * m + col + 1

		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1506/A
  Tags: math
  Rating: 800
*/
