//go:build codeforces_1698A

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
		first := 0
		for i := range n {
			a := nextInt()
			if i == 0 { first = a }
		}

		fmt.Fprintln(wr, first)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1698/A
  Tags: bitmasks, brute force
  Rating: 800
*/
