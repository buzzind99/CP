//go:build codeforces_50A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
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
	sol := math.Floor(float64(m*n)/2.0)

	fmt.Println(sol)
}

/*
  Link: https://codeforces.com/problemset/problem/50/A
  Tags: greedy, math
  Rating: 800
*/
