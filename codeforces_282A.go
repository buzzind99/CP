//go:build codeforces_282A

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
	sol := 0
    
	n := nextInt()
	for range n {
		str := next()
		if str[1] == '+' {
			sol++
		} else {
			sol--
		}
	}

	fmt.Println(sol)
}

/*
  Link: https://codeforces.com/problemset/problem/282/A
  Tags: implementation
  Rating: 800
*/
