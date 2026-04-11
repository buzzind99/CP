//go:build codeforces_2000A

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
		a := next()

		if a[:2] != "10" {
			fmt.Fprintln(wr, "NO")
		} else {
			if a[2] == '0' {
				fmt.Fprintln(wr, "NO")
			} else {
				fmt.Fprintln(wr, "YES")
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2000/A
  Tags: implementation, math, strings
  Rating: 800
*/
