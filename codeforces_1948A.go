//go:build codeforces_1948A

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

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		if !isEven(n) {
			fmt.Fprintln(wr, "NO")
		} else {
			fmt.Fprintln(wr, "YES")
			for i := 0; i < n/2; i++ {
				if i % 2 == 0 {
					fmt.Fprint(wr, "AA")
				} else {
					fmt.Fprint(wr, "BB")
				}
			}
			fmt.Fprintln(wr)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1948/A
  Tags: brute force, constructive algorithms
  Rating: 800
*/
