//go:build codeforces_432A

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

	n, k := nextInt(), nextInt()
	eligible := 0
    for range n {
        if 5 - nextInt() >= k { eligible++ }
    }

    fmt.Println(eligible/3)
}

/*
  Link: https://codeforces.com/problemset/problem/432/A
  Tags: greedy, implementation, sorting
  Rating: 800
*/
