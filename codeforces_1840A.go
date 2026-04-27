//go:build codeforces_1840A

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

func printSubarray(b []int, c []int) {
	fmt.Fprintln(wr, len(b), len(c))
	for _, v := range b {
		fmt.Fprint(wr, v, " ")
	}
	fmt.Fprintln(wr)
	for _, v := range c {
		fmt.Fprint(wr, v, " ")
	}
	fmt.Fprintln(wr)
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		s := next()
		ans := make([]byte, 0, n)
		for i := 0; i < n; {
			startChar := s[i]
			ans = append(ans, startChar)
			i++
			for i < n && s[i] != startChar { i++ }
			i++
		}
		
		fmt.Fprintln(wr, string(ans))
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1840/A
  Tags: implementation, strings, two pointers
  Rating: 800
*/
