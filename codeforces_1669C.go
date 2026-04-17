//go:build codeforces_1669C

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

func isEven(x int) bool {
	return x%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		modOdd := nextInt()%2
		modEven := nextInt()%2
		possible := true
		for i := range n-2 {
			a := nextInt()
			if isEven(i+1) {
				if a%2 != modEven { possible = false }
			} else {
				if a%2 != modOdd { possible = false }
			}
		}

		if possible {
			fmt.Fprintln(wr, "YES")
		} else {
			fmt.Fprintln(wr, "NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1669/C
  Tags: greedy, implementation, math
  Rating: 800
*/
