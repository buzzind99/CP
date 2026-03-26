//go:build codeforces_677A

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

    n, h := nextInt(), nextInt()
	count := 0
	for range n {
		curr := nextInt()
		if curr > h { count += 2 } else { count++ }
	}
	
	fmt.Println(count)
}

/*
  Link: https://codeforces.com/problemset/problem/677/A
  Tags: implementation
  Rating: 800
*/
