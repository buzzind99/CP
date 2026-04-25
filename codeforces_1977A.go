//go:build codeforces_1977A

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

func isEven (n int) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n, m := nextInt(), nextInt()

		if n >= m && isEven(n-m) {
			fmt.Fprintln(wr, "Yes")
		} else {
			fmt.Fprintln(wr, "No")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1977/A
  Tags: math
  Rating: 800
*/
