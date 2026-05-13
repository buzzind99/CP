//go:build codeforces_1932A

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

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n, s := nextInt(), next()
		count := 0
		for i := 0; i < n; {
			curr, next, nextNext := s[i], byte('.'), byte('.')
			if i+1 <= n-1 { next = s[i+1] }
			if i+2 <= n-1 { nextNext = s[i+2] }
			if curr == '@' { count++ }
			if next == '.' || next == '@' {
				i++
				continue
			} else if nextNext == '.' || nextNext == '@' {
				i += 2
				continue
			} else {
				break
			}
		}

		fmt.Fprintln(wr, count)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1932/A
  Tags: dp, greedy, implementation
  Rating: 800
*/
