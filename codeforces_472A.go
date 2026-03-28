//go:build codeforces_472A

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

	n := nextInt()
	if n%2 == 0 {
        fmt.Println(4, n-4)
    } else {
        fmt.Println(9, n-9)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/472/A
  Tags: math, number theory
  Rating: 800
*/
