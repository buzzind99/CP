//go:build codeforces_1996A

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
        n := nextInt()
		fmt.Println((n+2)/4)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1996/A
  Tags: binary search, math, ternary search
  Rating: 800
*/
