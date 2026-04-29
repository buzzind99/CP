//go:build codeforces_2126B

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
		n, k := nextInt(), nextInt()
		count, ans := 0, 0
		for range n {
			a := nextInt()
			if a == 0 {
				count++
			} else {
				ans += (count+1) /(k+1)
				count = 0
			}
		}
		ans += (count+1)/(k+1)
		count = 0

		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2126/B
  Tags: dp, greedy
  Rating: 800
*/
