//go:build codeforces_1866A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
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

	n := nextInt()
	minVal := math.MaxInt
	for range n {
		a := nextInt()
		minVal = min(minVal, abs(a))
	}

	fmt.Fprintln(wr, minVal)
}

/*
  Link: https://codeforces.com/problemset/problem/1866/A
  Tags: math
  Rating: 800
*/
