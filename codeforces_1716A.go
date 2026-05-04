//go:build codeforces_1716A

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

		if n == 1 {
			fmt.Fprintln(wr, 2)
		} else if n%3 == 0 {
			fmt.Fprintln(wr, n/3)
		} else {
			fmt.Fprintln(wr, n/3+1)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1716/A
  Tags: greedy, math
  Rating: 800
*/
