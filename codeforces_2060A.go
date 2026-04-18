//go:build codeforces_2060A

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
		a, b, d, e := nextInt(), nextInt(), nextInt(), nextInt()
		m := make(map[int]int)
		m[a+b]++
		m[d-b]++
		m[e-d]++
		maxCount := 0
		for _, v := range m {
			maxCount = max(maxCount, v)
		}

		fmt.Fprintln(wr, maxCount)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2060/A
  Tags: brute force
  Rating: 800
*/
