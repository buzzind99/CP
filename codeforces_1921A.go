//go:build codeforces_1921A

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
		minX, maxX := 1001, -1001
		minY, maxY := 1001, -1001
		for range 4 {
			x, y := nextInt(), nextInt()
			minX, maxX = min(minX, x), max(maxX, x)
			minY, maxY = min(minY, y), max(maxY, y)
		}

		fmt.Fprintln(wr, (maxX-minX)*(maxY-minY))
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1921/A
  Tags: greedy, math
  Rating: 800
*/
