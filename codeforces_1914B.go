//go:build codeforces_1914B

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
		n, k := nextInt(), nextInt()

		for i := n; i > k+1; i-- {
			fmt.Fprintf(wr, "%d ", i)
		}
		for i := 1; i <= k+1; i++ {
			fmt.Fprintf(wr, "%d ", i)
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1914/B
  Tags: constructive algorithms, math
  Rating: 800
*/
