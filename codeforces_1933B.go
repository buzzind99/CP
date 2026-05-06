//go:build codeforces_1933B

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

func abs(x int) int {
	if x < 0 { return -x }
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		sum := 0
		one := false
		for range n {
			a := nextInt()
			sum += a
			if a%3 == 1 { one = true }
		}

		sumMod := sum%3
		if sumMod == 0 {
			fmt.Fprintln(wr, 0)
		} else if sumMod == 2 {
			fmt.Fprintln(wr, 1)
		} else if sumMod == 1 {
			if one {
				fmt.Fprintln(wr, 1)
			} else {
				fmt.Fprintln(wr, 2)
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1933/B
  Tags: implementation, math, number theory
  Rating: 800
*/
