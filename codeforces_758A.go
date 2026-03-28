//go:build codeforces_758A

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
    
	n := nextInt()
	arr := make([]int, n)
	maxW := -1
	for i := range n {
		a := nextInt()
		arr[i] = a
		if a > maxW { maxW = a }
	}
	sum := 0
	for _, v := range arr {
		sum += abs(maxW-v)
	}

	fmt.Println(sum)
}

/*
  Link: https://codeforces.com/problemset/problem/758/A
  Tags: implementation, math
  Rating: 800
*/
