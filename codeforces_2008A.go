//go:build codeforces_2008A

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

		leftover := 2*(b%2)
		if leftover > a || (a-leftover)%2 != 0 {
			fmt.Fprintln(wr, "NO")
		} else {
			fmt.Fprintln(wr, "YES")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2008/A
  Tags: brute force, constructive algorithms, greedy, math
  Rating: 800
*/
