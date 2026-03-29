//go:build codeforces_432A

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

	n, k := nextInt(), nextInt()
	arr := make([]int, n)
	for i := range n {
		arr[i] = nextInt()
	}

	slices.Sort(arr)
	maxTeam := n/3
	count := 0
	for i := range maxTeam {
		if arr[(i+1)*3-1]+k <= 5 { count++ }
	}

	fmt.Println(count)
}

/*
  Link: https://codeforces.com/problemset/problem/432/A
  Tags: greedy, implementation, sorting
  Rating: 800
*/
