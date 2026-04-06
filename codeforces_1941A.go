//go:build codeforces_1941A

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
		n, m, k := nextInt(), nextInt(), nextInt()
		left, right := make([]int, 0, n), make([]int, 0, n)
		for range n {
			left = append(left, nextInt())
		}
		for range m {
			right = append(right, nextInt())
		}
		slices.Sort(left)
		slices.Sort(right)

		count := 0
		for i := range n {
			if left[i] >= k { break }
			for j := range m {
				if left[i]+right[j] > k { break }
				count++
			}
		}

		fmt.Fprintln(wr, count)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1941/A
  Tags: brute force, math
  Rating: 800
*/
