//go:build codeforces_141A

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func sortString(s string) string {
    r := []rune(s)
    slices.Sort(r)
    return string(r)
}

func main() {
	sc.Split(bufio.ScanWords)

	a, b, c := next(), next(), sortString(next())
	if c == sortString(a+b) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/141/A
  Tags: implementation, sortings, strings
  Rating: 800
*/
