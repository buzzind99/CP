//go:build codeforces_1560B

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
		a, b, c := nextInt(), nextInt(), nextInt()
		diff := abs(a-b)
		maxVal := diff*2

		if maxVal < a || maxVal < b || maxVal < c {
			fmt.Fprintln(wr, -1)
		} else {
			if c <= diff {
				fmt.Fprintln(wr, c+diff)
			} else {
				fmt.Fprintln(wr, c-diff)
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1560/B
  Tags: math
  Rating: 800
*/
