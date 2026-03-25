//go:build codeforces_112A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	s1 := strings.ToLower(next())
	s2 := strings.ToLower(next())

	if s1 < s2 {
		fmt.Println(-1)
	} else if s1 == s2 {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/112/A
  Tags: implementation, strings
  Rating: 800
*/
