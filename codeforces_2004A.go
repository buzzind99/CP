//go:build codeforces_2004A

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

func abs(x int) int {
	if x < 0 { return -x }
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		if n == 2 {
			a, b := nextInt(), nextInt()
			diff := abs(a-b)
			if diff > 1 {
				fmt.Fprintln(wr, "YES")
			} else {
				fmt.Fprintln(wr, "NO")
			}

			continue
		}

		for range n {
			nextInt()
		}

		fmt.Fprintln(wr, "NO")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2004/A
  Tags: implementation, math
  Rating: 800
*/
