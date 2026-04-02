//go:build codeforces_1360B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
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

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		slice := make([]int, 0, n)
		for range n {
			slice = append(slice, nextInt())
		}

		slices.Sort(slice)
		minDiff := math.MaxInt
		for i := range len(slice)-1 {
			minDiff = min(slice[i+1]-slice[i], minDiff)
		}

		fmt.Fprintln(wr, minDiff)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1360/B
  Tags: greedy, sortings
  Rating: 800
*/
