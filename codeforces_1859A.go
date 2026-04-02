//go:build codeforces_1859A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
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

func printSubarray(b []int, c []int) {
	fmt.Fprintln(wr, len(b), len(c))
	for _, v := range b {
		fmt.Fprint(wr, v, " ")
	}
	fmt.Fprintln(wr)
	for _, v := range c {
		fmt.Fprint(wr, v, " ")
	}
	fmt.Fprintln(wr)
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		a := make([]int, 0, n)
		minVal := math.MaxInt
		for range n {
			val := nextInt()
			a = append(a, val)
			minVal = min(val, minVal)
		}

		b, c := make([]int, 0, n-1), make([]int, 0, n-1)
		for _, v := range a {
			if v == minVal {
				b = append(b, v)
			} else {
				c = append(c, v)
			}
		}

		if len(c) == 0 {
			fmt.Fprintln(wr, -1)
		} else {
			printSubarray(b, c)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1859/A
  Tags: constructive algorithms, math, number theory
  Rating: 800
*/
