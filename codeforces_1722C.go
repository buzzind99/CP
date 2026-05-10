//go:build codeforces_1722C

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
		m := make(map[string][]int)
		for player := 1; player <= 3; player++ {
			for range n {
				s := next()
				m[s] = append(m[s], player)
			}
		}

		a, b, c := 0, 0, 0
		for _, val := range m {
			l := len(val)
			multiplier := 3
			if l == 2 { multiplier = 1 } else if l == 3 { multiplier = 0 }
			for _, player := range val {
				if player == 1 { a += multiplier } else if player == 2 { b += multiplier } else { c += multiplier }
			}
		}

		fmt.Fprintln(wr, a, b, c)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1722/C
  Tags: data structures, implementation
  Rating: 800
*/
