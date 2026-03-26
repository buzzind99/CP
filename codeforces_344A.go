//go:build codeforces_344A

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
	prev := ""
	for range n {
		s := next()
		if prev != s {
			count++
			prev = s
		}
	}

	fmt.Println(count)
}

/*
  Link: https://codeforces.com/problemset/problem/344/A
  Tags: implementation
  Rating: 800
*/
