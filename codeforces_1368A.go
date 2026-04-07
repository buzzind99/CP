//go:build codeforces_1368A

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
		a, b, n := nextInt(), nextInt(), nextInt()

		count := 0
		for a <= n && b <= n {
			if a < b { a += b } else { b += a }
			count++
		}

		fmt.Fprintln(wr, count)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1368/A
  Tags: brute force, greedy, implementation, math
  Rating: 800
*/
