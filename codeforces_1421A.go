//go:build codeforces_1421A

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

	n := nextInt()
	for range n {
		a, b := nextInt(), nextInt()

		fmt.Fprintln(wr, a^b)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1421/A
  Tags: bitmasks, greedy, math
  Rating: 800
*/
