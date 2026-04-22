//go:build codeforces_1538A

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

func isOdd(x int) bool {
	return x%2 != 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		a := make([]int, n)
		minVal, maxVal := 101, -1
		minIdx, maxIdx := 0, 0
		for i := 0; i < n; i++ {
			a[i] = nextInt()
			if a[i] < minVal {
				minVal, minIdx = a[i], i
			}
			if a[i] > maxVal {
				maxVal, maxIdx = a[i], i
			}
		}

		p1 := minIdx
		p2 := maxIdx
		if p1 > p2 {
			p1, p2 = p2, p1
		}
		ans := p2 + 1
		ans = min(ans, n - p1)
		ans = min(ans, (p1 + 1) + (n - p2))

		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1538/A
  Tags: brute force, dp, greedy
  Rating: 800
*/
