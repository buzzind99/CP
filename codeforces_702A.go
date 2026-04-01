//go:build codeforces_702A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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

	n := nextInt()
	maxL, prev, count := 0, 0, 0
	for range n {
		a := nextInt()
		if a > prev {
			count++
		} else {
			if count > maxL { maxL = count }
			count = 1
		}
		prev = a
	}
	if count > maxL { maxL = count }

	fmt.Println(maxL)
}

/*
  Link: https://codeforces.com/problemset/problem/702/A
  Tags: dp, greedy, implementation
  Rating: 800
*/
