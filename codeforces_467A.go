//go:build codeforces_467A

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
	count := 0
	for range n {
		pi, qi := nextInt(), nextInt()
		if pi+2 <= qi { count++ }
	}

	fmt.Println(count)
}

/*
  Link: https://codeforces.com/problemset/problem/467/A
  Tags: implementation
  Rating: 800
*/
