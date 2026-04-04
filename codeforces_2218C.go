//go:build codeforces_2218C

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
			second := 3*n-2*i+1
			largest := 3*n-2*i+2
			fmt.Fprintf(wr, "%d %d %d ", i, second, largest)
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2218/C
  Tags: -
  Rating: -
  Contest: Codeforces Round 1090 (Div. 4)
*/
