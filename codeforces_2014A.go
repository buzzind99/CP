//go:build codeforces_2014A

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
		n, k := nextInt(), nextInt()
		coins, count := 0, 0
		for range n {
			a := nextInt()
			if a == 0 && coins > 0 { coins--; count++ } else if a >= k { coins += a }
		}

		fmt.Fprintln(wr, count)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2014/A
  Tags: greedy, implementation
  Rating: 800
*/
