//go:build codeforces_2009B

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
		n := nextInt()
		arr := make([]int, 0, n)
		for range n {
			s := next()
			for i := range len(s) {
				if s[i] == '#' {
					arr = append(arr, i+1)
				}
			}
		}

		for i := len(arr)-1; i >= 0; i-- {
			fmt.Fprint(wr, arr[i], " ")
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2009/B
  Tags: brute force, implementation
  Rating: 800
*/
