//go:build codeforces_1030A

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
	for range n {
		x := nextInt()
		if x == 1 {
			fmt.Println("HARD")
			return
		}
	}
	fmt.Println("EASY")
}

/*
  Link: https://codeforces.com/problemset/problem/1030/A
  Tags: implementation
  Rating: 800
*/
