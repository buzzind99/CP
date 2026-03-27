//go:build codeforces_1154A

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

	arr := []int{nextInt(), nextInt(), nextInt(), nextInt()}
	slices.Sort(arr)
	
	fmt.Printf("%d %d %d", arr[3]-arr[0], arr[3]-arr[1], arr[3]-arr[2])
}

/*
  Link: https://codeforces.com/problemset/problem/1154/A
  Tags: math
  Rating: 800
*/
