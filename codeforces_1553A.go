//go:build codeforces_1553A

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

		fmt.Fprintln(wr, (n+1)/10)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1553/A
  Tags: math, number theory
  Rating: 800
*/
