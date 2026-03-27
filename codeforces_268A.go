//go:build codeforces_268A

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
	team := make([][]int, n)
	for i := range team {
		team[i] = []int{nextInt(), nextInt()}
	}

	count := 0
	for i := range n {
		for j := range n {
			if i == j { continue }
			if team[i][0] == team[j][1] { count++ }
		}
	}
    
	fmt.Println(count)
}

/*
  Link: https://codeforces.com/problemset/problem/268/A
  Tags: brute force
  Rating: 800
*/
