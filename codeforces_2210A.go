//go:build codeforces_2210A

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
        n := nextInt()
		for i := 0; i < n; i++ {
			fmt.Print(n-i, " ")
		}
		fmt.Println()
    }
}

/*
  Link: https://codeforces.com/problemset/problem/2210/A
  Tags: constructive algorithms, dp, greedy, number theory
  Rating: -
*/
