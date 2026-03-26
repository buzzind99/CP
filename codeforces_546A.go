//go:build codeforces_546A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func seqSum(n int) int { return n*(n+1)/2 }

func max(n int, m int) int {
	if n < m { return m }
	return n
}

func main() {
	sc.Split(bufio.ScanWords)

    k, n, w := nextInt(), nextInt(), nextInt()

	fmt.Println(max(k*seqSum((w))-n, 0))
}

/*
  Link: https://codeforces.com/problemset/problem/546/A
  Tags: brute force, implementation, math
  Rating: 800
*/
