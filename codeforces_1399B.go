//go:build codeforces_1399B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
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
		candies, oranges := make([]int, 0, n), make([]int, 0, n)
		for range n {
			candies = append(candies, nextInt())
		}
		for range n {
			oranges = append(oranges, nextInt())
		}

		targetCandies, targetOranges := slices.Min(candies), slices.Min(oranges)
		sum := 0
		for i := range n {
			sum += max(candies[i]-targetCandies, oranges[i]-targetOranges)
		}

		fmt.Fprintln(wr, sum)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1399/B
  Tags: greedy
  Rating: 800
*/
