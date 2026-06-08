//go:build codeforces_2148B

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
		n, m, x, y := nextInt(), nextInt(), nextInt(), nextInt()
		hCross := 0
		for range n {
			a := nextInt()
			if a > 0 && a < y {
				hCross++
			}
		}
		vCross := 0
		for range m {
			b := nextInt()
			if b > 0 && b < x {
				vCross++
			}
		}

		fmt.Fprintln(wr, hCross+vCross)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2148/B
  Tags: geometry
  Rating: 800
*/
