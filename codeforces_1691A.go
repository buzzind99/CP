//go:build codeforces_1691A

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
		odd, even := 0, 0
		for range n {
			a := nextInt()
			if isEven(a) { even++ } else { odd++ }
		}

		if even > odd {
			fmt.Fprintln(wr, odd)
		} else {
			fmt.Fprintln(wr, even)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1691/A
  Tags: brute force, greedy, math
  Rating: 800
*/
