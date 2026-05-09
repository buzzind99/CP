//go:build codeforces_1833A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

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
	defer wr.Flush()

	t := nextInt()
	for range t {
		n, s := nextInt(), next()
		exists := [7][7]bool{}
		count := 0
		for i := 0; i < n-1; i++ {
			u, v := s[i]-'a', s[i+1]-'a'
			if !exists[u][v] {
				exists[u][v] = true
				count++
			}
		}

		fmt.Fprintln(wr, count)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1833/A
  Tags: implementation, strings
  Rating: 800
*/
