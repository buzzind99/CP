//go:build codeforces_1950B

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
		colCount := 1
		left := true
		for range 2*n {
			rowCount := 1
			dot := !left
			for range 2*n {
				if !dot {
					fmt.Fprint(wr, "#")
					if rowCount == 2 { rowCount = 1; dot = !dot } else { rowCount++ }
				} else {
					fmt.Fprint(wr, ".")
					if rowCount == 2 { rowCount = 1; dot = !dot } else { rowCount++ }
				}
			}
			if colCount == 2 { colCount = 1; left = !left } else { colCount++ }
			fmt.Fprintln(wr)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1950/B
  Tags: implementation
  Rating: 800
*/
