//go:build codeforces_2074B

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
		sum := 0
		for range n {
			sum += nextInt()
		}

		fmt.Fprintln(wr, sum-n+1)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2074/B
  Tags: geometry, greedy, math
  Rating: 800
*/
