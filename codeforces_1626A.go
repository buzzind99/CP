//go:build codeforces_1626A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
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
		s := next()
		bytes := []byte(s)
		slices.Sort(bytes)

		fmt.Fprintf(wr, "%s\n", bytes)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1626/A
  Tags: constructive algorithms, sortings
  Rating: 800
*/
