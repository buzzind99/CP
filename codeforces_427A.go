//go:build codeforces_427A

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
	police := 0
	count := 0
	for range n {
		x := nextInt()
		police += x
		if police < 0 {
			count++
			police = 0
		}
	}

	fmt.Println(count)
}

/*
  Link: https://codeforces.com/problemset/problem/427/A
  Tags: implementation
  Rating: 800
*/
