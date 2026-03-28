//go:build codeforces_1901A

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

	t := nextInt()
    for range t {
        n, x := nextInt(), nextInt()
		maxV := 0
		last := 0
		for range n {
			a := nextInt()
			dist := a-last
			maxV = max(maxV, dist)
			last = a
		}
		dist := (x-last)*2
		maxV = max(maxV, dist)

		fmt.Println(maxV)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1901/A
  Tags: greedy, math
  Rating: 800
*/
