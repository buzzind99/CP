//go:build codeforces_1855A

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
		count := 0
		for i := 1; i <= n; i++ {
			a := nextInt()
			if a == i { count++ }
		}

		fmt.Fprintln(wr, (count+1)/2)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1855/A
  Tags: greedy, math
  Rating: 800
*/
