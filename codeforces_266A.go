//go:build codeforces_266A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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

	nextInt()
	s := next()
	sol := 0
	var prev rune
	for _, curr := range s {
		if curr == prev {
			sol++
			continue
		}
		prev = curr
	}
    
	fmt.Println(sol)
}

/*
  Link: https://codeforces.com/problemset/problem/266/A
  Tags: implementation
  Rating: 800
*/
