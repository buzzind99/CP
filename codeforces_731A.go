//go:build codeforces_731A

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

	s := next()
	curr := byte('a')
	sum := 0
	for i := range s {
		diff := abs(int(s[i]-'a')-int(curr-'a'))
		sum += min (diff, 26-diff)
		curr = s[i]
	}

	fmt.Fprintln(wr, sum)
}

/*
  Link: https://codeforces.com/problemset/problem/731/A
  Tags: implementation, strings
  Rating: 800
*/
