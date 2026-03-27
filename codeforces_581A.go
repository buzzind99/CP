//go:build codeforces_581A

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

    r, b := nextInt(), nextInt()
	minRb := min(r, b)
	maxRb := max(r, b)

	fmt.Print(minRb, (maxRb-minRb)/2)
}

/*
  Link: https://codeforces.com/problemset/problem/581/A
  Tags: implementation, math
  Rating: 800
*/
