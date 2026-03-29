//go:build codeforces_1950A

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
		a, b, c := nextInt(), nextInt(), nextInt()
		if a < b && b < c {
			fmt.Println("STAIR")
		} else if a < b && b > c {
			fmt.Println("PEAK")
		} else {
			fmt.Println("NONE")
		}
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1950/A
  Tags: implementation
  Rating: 800
*/
