//go:build codeforces_723A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
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
    
	arr := []int{nextInt(), nextInt(), nextInt()}
	slices.Sort(arr)

	fmt.Print(arr[2] - arr[0])
}

/*
  Link: https://codeforces.com/problemset/problem/723/A
  Tags: implementation, math, sortings
  Rating: 800
*/
