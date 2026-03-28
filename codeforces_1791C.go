//go:build codeforces_1791C

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

	t := nextInt()
    for range t {
        n, s := nextInt(), next()
        l, r := 0, n-1
        for s[l] != s[r] && l < r {
            l++
            r--
        }
        fmt.Println(r-l+1)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1791/C
  Tags: implementation, two pointers
  Rating: 800
*/
