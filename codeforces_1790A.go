//go:build codeforces_1790A

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
	pi := "3141592653589793238462643383279"
	for range t {
		s := next()
		count := 0
		for i := range s {
			if s[i] != pi[i] { break }
			count++
		}

		fmt.Fprintln(wr, count)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1790/A
  Tags: implementation, math, strings
  Rating: 800
*/
