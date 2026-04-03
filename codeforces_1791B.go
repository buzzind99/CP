//go:build codeforces_1791B

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
		s := next()
		x, y := 0, 0
		candyFound := false
		for i := range n {
			switch s[i] {
			case 'U': y++
			case 'D': y--
			case 'R': x++
			case 'L': x--
			}
			if x == 1 && y == 1 {
				candyFound = true
				break
			}
		}
		if candyFound {
			fmt.Fprintln(wr, "YES")
		} else {
			fmt.Fprintln(wr, "NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1791/B
  Tags: geometry, implementation
  Rating: 800
*/
