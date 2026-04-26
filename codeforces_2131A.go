//go:build codeforces_2131A

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
		a, b := make([]int, 0, n), make([]int, 0, n)
		for range n {
			a = append(a, nextInt())
		}
		for range n {
			b = append(b, nextInt())
		}

		ops := 1
		for i := range a {
			if a[i] > b[i] { ops += a[i]-b[i] }
		}

		fmt.Fprintln(wr, ops)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2131/A
  Tags: math
  Rating: 800
*/
