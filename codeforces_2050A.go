//go:build codeforces_2050A

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
		n, m := nextInt(), nextInt()
		currentLen, wordCount := 0, 0
		canFit := true
		for i := 0; i < n; i++ {
			s := next()
			if canFit {
				if currentLen + len(s) <= m {
					currentLen += len(s)
					wordCount++
				} else {
					canFit = false
				}
			}
		}

		fmt.Fprintln(wr, wordCount)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2050/A
  Tags: implementation
  Rating: 800
*/
