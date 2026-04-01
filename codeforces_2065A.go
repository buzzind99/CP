//go:build codeforces_2065A

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
	t := nextInt()
	for range t {
		s := next()
		fmt.Printf("%si\n", s[:len(s)-2])
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2065/A
  Tags: brute force, constructive algorithms, greedy, implementation, strings
  Rating: 800
*/
