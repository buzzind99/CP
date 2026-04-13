//go:build codeforces_1872A

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
		a, b, c := nextInt(), nextInt(), nextInt()
		if a == b {
			fmt.Fprintln(wr, 0)
		} else {
			r := abs(b-a)
			count := 0
			for r > 2*c {
				if a < b {
					a += c
					b -= c
				} else {
					a -= c
					b += c
				}
				count++
				r = abs(b-a)
			}
	
			fmt.Fprintln(wr, count+1)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1872/A
  Tags: brute force, greedy, math
  Rating: 800
*/
