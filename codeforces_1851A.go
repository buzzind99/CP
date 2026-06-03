//go:build codeforces_1851A

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
		n, m, k, H := nextInt(), nextInt(), nextInt(), nextInt()
		hMap := make(map[int]struct{})
		for i := 1; i < m; i++ {
			keyPlus, keyMinus := H+k*i, H-k*i
			hMap[keyPlus], hMap[keyMinus] = struct{}{}, struct{}{}
		}

		count := 0
		for range n {
			if _, exists := hMap[nextInt()]; exists { count++ }
		}

		fmt.Fprintln(wr, count)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1851/A
  Tags: brute force, constructive algorithms, math
  Rating: 800
*/
