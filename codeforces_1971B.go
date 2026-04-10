//go:build codeforces_1971B

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
		s := []rune(next())
		unique := true
		prev := s[0]
		for i := 1; i < len(s); i++ {
			if s[i] != prev { unique = false }
			prev = s[i]
		}

		if unique {
			fmt.Fprintln(wr, "NO")
			continue
		}

		prev = s[0]
		for i := 1; i < len(s); i++ {
			if s[i] != prev { s[i], s[i-1] = s[i-1], s[i]; break }
			prev = s[i]
		}

		fmt.Fprintln(wr, "YES")
		fmt.Fprintln(wr, string(s))
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1971/B
  Tags: implementation, strings
  Rating: 800
*/
