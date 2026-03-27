//go:build codeforces_1409A

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

func abs(x int) int {
    if x < 0 { return -x }
    return x
}

func main() {
	sc.Split(bufio.ScanWords)

	t := nextInt()
	for range t {
        a, b := nextInt(), nextInt()
        if a == b {
            fmt.Println(0)
        } else {
            fmt.Println((abs(a-b)+9)/10)
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1409/A
  Tags: greedy, math
  Rating: 800
*/
