//go:build codeforces_2131B

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
		for i := range n {
			if isEven(i) {
				fmt.Fprint(wr, -1, " ")
			} else {
				if i+1 < n {
					fmt.Fprint(wr, 3, " ")
				} else {
					fmt.Fprint(wr, 2, " ")
				}
			}
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2131/B
  Tags: constructive algorithms, greedy, math
  Rating: 800
*/
