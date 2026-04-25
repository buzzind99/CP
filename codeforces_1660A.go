//go:build codeforces_1660A

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
		a, b := nextInt(), nextInt()

		if a == 0 {
			fmt.Fprintln(wr, 1)
		} else {
			sum := a*1+b*2

			fmt.Fprintln(wr, sum+1)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1660/A
  Tags: greedy, math
  Rating: 800
*/
