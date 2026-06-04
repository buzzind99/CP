//go:build codeforces_2117B

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
		for i := 1; i <= n; i++ {
			if i == 1 {
				fmt.Fprint(wr, 2, " ")
			} else if i == n {
				fmt.Fprint(wr, 1, " ")
			} else {
				fmt.Fprint(wr, i+1, " ")
			}
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2117/B
  Tags: constructive algorithms
  Rating: 800
*/
