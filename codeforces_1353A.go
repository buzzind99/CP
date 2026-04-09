//go:build codeforces_1353A

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
		n, m := nextInt(), nextInt()
		if n == 1 {
			fmt.Fprintln(wr, 0)
		} else if n == 2 {
			fmt.Fprintln(wr, m)
		} else {
			fmt.Fprintln(wr, 2*m)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1353/A
  Tags: constructive algorithms, greedy, math
  Rating: 800
*/
