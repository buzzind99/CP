//go:build codeforces_750A

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
    
	n, t := nextInt(), nextInt()
	sol := 0
	for i := n; i > 0; i-- {
		nt := 5*i*(i+1)/2 + t
		if nt <= 240 {
			sol = i
			break
		}
	}

	fmt.Print(sol)
}

/*
  Link: https://codeforces.com/problemset/problem/750/A
  Tags: binary search, brute force, implementation, math
  Rating: 800
*/
