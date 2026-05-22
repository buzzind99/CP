//go:build codeforces_1631A

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
		A, B := make([]int, 0, n), make([]int, 0, n)
		for range n {
			A = append(A, nextInt())
		}
		for range n {
			B = append(B, nextInt())
		}

		for i := range n {
			if B[i] > A[i] { A[i], B[i] = B[i], A[i] }
		}
		maxA, maxB := slices.Max(A), slices.Max(B)

		fmt.Fprintln(wr, maxA*maxB)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1631/A
  Tags: greedy
  Rating: 800
*/
