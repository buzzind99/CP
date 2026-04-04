//go:build codeforces_2218E

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

func maxXor(nums []int) int {
	maxResult, mask := 0, 0
	for i := 31; i >= 0; i-- {
		mask |= (1 << i)
		prefixes := make(map[int]bool)
		for _, num := range nums {
			prefixes[num&mask] = true
		}
		tempTarget := maxResult | (1 << i)

		for prefix := range prefixes {
			if prefixes[prefix^tempTarget] {
				maxResult = tempTarget
				break
			}
		}
	}
	return maxResult
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		arr := make([]int, 0, n)
		for range n {
			arr = append(arr, nextInt())
		}
		res := maxXor(arr)

		fmt.Fprintln(wr, res)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2218/E
  Tags: -
  Rating: -
  Contest: Codeforces Round 1090 (Div. 4)
*/
