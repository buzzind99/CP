//go:build codeforces_2093A

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

func isEven(x int) bool {
	return x%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		k := nextInt()

		if isEven(k) {
			fmt.Fprintln(wr, "NO")
		} else {
			fmt.Fprintln(wr, "YES")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2093/A
  Tags: math
  Rating: 800
*/
