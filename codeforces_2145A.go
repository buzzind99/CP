//go:build codeforces_2145A

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
		mod := n%3

		if mod == 0 {
			fmt.Fprintln(wr, 0)
		} else if mod == 1 {
			fmt.Fprintln(wr, 2)
		} else if mod == 2 {
			fmt.Fprintln(wr, 1)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2145/A
  Tags: math
  Rating: 800
*/
