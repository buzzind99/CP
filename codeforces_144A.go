//go:build codeforces_144A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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

	n := nextInt()
	min, max := 101, -1
	minIdx, maxIdx := 0, 0
	for i := range n {
		a := nextInt()
		if a > max { max, maxIdx = a, i }
		if a <= min { min, minIdx = a, i }
	}

	sol := maxIdx + (n - 1 - minIdx)
    if maxIdx > minIdx { sol-- }

	fmt.Println(sol)
}

/*
  Link: https://codeforces.com/problemset/problem/144/A
  Tags: implementation
  Rating: 800
*/
