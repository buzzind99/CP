//go:build codeforces_1676C

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

	t := nextInt()
	for range t {
		n, m := nextInt(), nextInt()
		arr := make([]string, 0, n)
		for range n {
			s := next()
			arr = append(arr, s)
		}

		minCost := math.MaxInt32
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				currentCost := 0
				for k := 0; k < m; k++ {
					currentCost += abs(int(arr[i][k]) - int(arr[j][k]))
				}

				if currentCost < minCost {
					minCost = currentCost
				}
			}
		}

		fmt.Println(minCost)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1676/C
  Tags: brute force, greedy, implementation, math, strings
  Rating: 800
*/
