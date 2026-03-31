//go:build codeforces_749A

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
	fmt.Println(n/2)
	if n%2 == 0 {
		for range n/2 { fmt.Print(2, " ") }
	} else {
		fmt.Print(3, " ")
		for range n/2-1 { fmt.Print(2, " ") }
	}
}

/*
  Link: https://codeforces.com/problemset/problem/749/A
  Tags: greedy, implementation, math, number theory
  Rating: 800
*/
