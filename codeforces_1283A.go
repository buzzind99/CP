//go:build codeforces_1283A

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

	t := nextInt()
	for range t {
		h, m := nextInt(), nextInt()
		sol := (23-h)*60+(60-m)

		fmt.Println(sol)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1283/A
  Tags: math
  Rating: 800
*/
