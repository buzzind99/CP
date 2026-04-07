//go:build codeforces_1311A

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

		if a == b {
			fmt.Fprintln(wr, 0)
		} else if a < b {
			if (b-a)%2 == 0 {
				fmt.Fprintln(wr, 2)
			} else {
				fmt.Fprintln(wr, 1)
			}
		} else {
			if (b-a)%2 == 0 {
				fmt.Fprintln(wr, 1)
			} else {
				fmt.Fprintln(wr, 2)
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1311/A
  Tags: greedy, implementation, math
  Rating: 800
*/
