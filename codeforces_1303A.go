//go:build codeforces_1303A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		s := next()
		first := strings.Index(s, "1")
		last := strings.LastIndex(s, "1")
		if first == -1 {
			fmt.Fprintln(wr, 0)
			continue
		}
		sub := s[first:last+1]
		zeroCount := strings.Count(sub, "0")

		fmt.Fprintln(wr, zeroCount)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1303/A
  Tags: implementation, strings
  Rating: 800
*/
