//go:build codeforces_2008C

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

func isEven(x int) bool {
	return x%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		l, r := nextInt(), nextInt()
		diff := r-l
		sum, count := 0, 1
		for sum < diff {
			sum += count
			count++
		}
		if sum > diff { count-- }

		fmt.Fprintln(wr, count)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2008/C
  Tags: binary search, brute force, math
  Rating: 800
*/
