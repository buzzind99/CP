//go:build codeforces_1669B

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
		m := make(map[int]int)
		sol := -1
		for range n {
			a := nextInt()
			if sol == -1 {
				m[a]++
				if m[a] == 3 { sol = a }
			}
		}

		fmt.Println(sol)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1669/B
  Tags: implementation, sortings
  Rating: 800
*/
