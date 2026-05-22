//go:build codeforces_1709A

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

func abs(x int) int {
    if x < 0 { return -x }
    return x
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		firstKey, a, b, c := nextInt(), nextInt(), nextInt(), nextInt()
		doors := [4]int{0,a,b,c}

		secondKey := doors[firstKey]
		if secondKey == 0 {
			fmt.Fprintln(wr, "NO")
			continue
		}
		thirdKey := doors[secondKey]
		if thirdKey == 0 {
			fmt.Fprintln(wr, "NO")
			continue
		}
		fmt.Fprintln(wr, "YES")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1709/A
  Tags: brute force, greedy, implementation, math
  Rating: 800
*/
