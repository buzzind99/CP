//go:build codeforces_116A

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
	curr := 0
	max := 0
	for range n {
		out, in := nextInt(), nextInt()
		curr -= out
		curr += in
		if curr > max { max = curr }
	}

	fmt.Println(max)
}

/*
  Link: https://codeforces.com/problemset/problem/116/A
  Tags: implementation
  Rating: 800
*/
