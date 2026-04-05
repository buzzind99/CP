//go:build codeforces_1992A

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
		a, b, c := nextInt(), nextInt(), nextInt()
		for range 5 {
			if a <= b && a <= c {
				a++
			} else if b <= a && b <= c {
				b++
			} else if c <= a && c <= b {
				c++
			}
		}

		fmt.Fprintln(wr, a*b*c)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1992/A
  Tags: brute force, constructive algorithms, greedy, math, sortings
  Rating: 800
*/
