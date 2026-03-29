//go:build codeforces_1374A

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
        x, y, n := nextInt(), nextInt(), nextInt()
		sol := (n/x)*x+y
        if sol > n { sol -= x}

		fmt.Println(sol)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1374/A
  Tags: math
  Rating: 800
*/
