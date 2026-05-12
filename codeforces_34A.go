//go:build codeforces_34A

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

	n := nextInt()
	arr := make([]int, 0, n)
	for range n {
		arr = append(arr, nextInt())
	}

	minDiff, idx1, idx2 := 1001, -1, -1
	for i := 0; i < n; i++ {
		nextIdx := (i+1)%n
		diff := arr[i] - arr[nextIdx]
		if diff < 0 { diff = -diff }

		if diff < minDiff {
			minDiff = diff
			idx1, idx2 = i+1, nextIdx+1
		}
	}

	fmt.Fprintln(wr, idx1, idx2)
}

/*
  Link: https://codeforces.com/problemset/problem/34/A
  Tags: implementation
  Rating: 800
*/
