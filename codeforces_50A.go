//go:build codeforces_50A

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
    
    m, n := nextInt(), nextInt()

    fmt.Println((m * n) / 2)
}

/*
  Link: https://codeforces.com/problemset/problem/50/A
  Tags: greedy, math
  Rating: 800
*/
