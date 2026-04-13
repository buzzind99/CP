//go:build codeforces_1358A

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

		fmt.Fprintln(wr, ((m*n)+1)/2)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1358/A
  Tags: greedy, math
  Rating: 800
*/
