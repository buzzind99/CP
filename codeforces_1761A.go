//go:build codeforces_1761A

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
		n, a, b := nextInt(), nextInt(), nextInt()
		if a == n && b == n {
			fmt.Fprintln(wr, "Yes")
		} else if a+b >= n-1 {
			fmt.Fprintln(wr, "No")
		} else if a+b <= n-2 {
			fmt.Fprintln(wr, "Yes")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1761/A
  Tags: brute force, constructive algorithms
  Rating: 800
*/
