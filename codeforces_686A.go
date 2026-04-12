//go:build codeforces_686A

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

	n := nextInt()
	z := nextInt()
	sum := z
	distressed := 0
	for range n {
		a, x := next(), nextInt()
		if a[0] == '+' {
			sum += x
		} else {
			if sum < x { distressed++; continue } else { sum -= x }
		}

	}

	fmt.Fprintln(wr, sum, distressed)
}

/*
  Link: https://codeforces.com/problemset/problem/686/A
  Tags: constructive algorithms, implementation
  Rating: 800
*/
