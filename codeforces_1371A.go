//go:build codeforces_1371A

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

		if n%2 == 0 {
			fmt.Fprintln(wr, n/2)
		} else {
			fmt.Fprintln(wr, n/2+1)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1371/A
  Tags: math
  Rating: 800
*/
