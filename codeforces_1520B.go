//go:build codeforces_1520B

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
		count := 0
		for powerOfTen := 1; powerOfTen <= 1000000000; powerOfTen = powerOfTen*10+1 {
			for digit := 1; digit <= 9; digit++ {
				if powerOfTen*digit <= n {
					count++
				}
			}
		}

		fmt.Fprintln(wr, count)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1520/B
  Tags: brute force, math, number theory
  Rating: 800
*/
