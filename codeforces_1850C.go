//go:build codeforces_1850C

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
		for range 8 {
			row := next()
			for _, char := range row {
				if char != '.' {
					fmt.Fprint(wr, string(char))
				}
			}
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1850/C
  Tags: implementation, strings
  Rating: 800
*/
