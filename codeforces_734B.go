//go:build codeforces_734B

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

	k2, k3, k5, k6 := nextInt(), nextInt(), nextInt(), nextInt()
	minBig := min(k2, k5, k6)
	minSmall := min(k3, k2-minBig)

	fmt.Fprintln(wr, minBig*256+minSmall*32)
}

/*
  Link: https://codeforces.com/problemset/problem/734/B
  Tags: brute force, greedy, implementation, math
  Rating: 800
*/
