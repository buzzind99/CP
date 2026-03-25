//go:build codeforces_791A

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

    a, b := nextInt(), nextInt()
	sol := 0
	for a <= b {
		a *= 3
		b *= 2
		sol++
	}

	fmt.Println(sol)
}

/*
  Link: https://codeforces.com/problemset/problem/791/A
  Tags: implementation
  Rating: 800
*/
