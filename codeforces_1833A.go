//go:build codeforces_1833A

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
		n, s := nextInt(), next()
		m := make(map[string]struct{})
		for i := range n-1 {
			notes := s[i:i+2]
			m[notes] = struct{}{}
		}

		fmt.Fprintln(wr, len(m))
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1833/A
  Tags: implementation, strings
  Rating: 800
*/
